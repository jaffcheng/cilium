// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sTest

import (
	"context"
	"fmt"
	"net"
	"time"

	. "github.com/cilium/cilium/test/ginkgo-ext"
	"github.com/cilium/cilium/test/helpers"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
)

var _ = Describe("K8sServicesTest", func() {
	var (
		kubectl                *helpers.Kubectl
		serviceName                               = "app1-service"
		backgroundCancel       context.CancelFunc = func() { return }
		backgroundError        error
		enableBackgroundReport = true
		ciliumPodK8s1          string
		testDSClient           = "zgroup=testDSClient"
		testDS                 = "zgroup=testDS"
		echoServiceName        = "echo"
		echoPodLabel           = "name=echo"
	)

	applyPolicy := func(path string) {
		By(fmt.Sprintf("Applying policy %s", path))
		_, err := kubectl.CiliumPolicyAction(helpers.DefaultNamespace, path, helpers.KubectlApply, helpers.HelperTimeout)
		ExpectWithOffset(1, err).Should(BeNil(), fmt.Sprintf("Error creating resource %s: %s", path, err))
	}

	BeforeAll(func() {
		var err error

		kubectl = helpers.CreateKubectl(helpers.K8s1VMName(), logger)
		DeployCiliumAndDNS(kubectl)

		ciliumPodK8s1, err = kubectl.GetCiliumPodOnNodeWithLabel(helpers.KubeSystemNamespace, helpers.K8s1)
		Expect(err).Should(BeNil(), "Cannot get cilium pod on k8s1")
	})

	AfterFailed(func() {
		kubectl.CiliumReport(helpers.KubeSystemNamespace,
			"cilium service list",
			"cilium endpoint list")
	})

	JustBeforeEach(func() {
		if enableBackgroundReport {
			backgroundCancel, backgroundError = kubectl.BackgroundReport("uptime")
			Expect(backgroundError).To(BeNil(), "Cannot start background report process")
		}
	})

	JustAfterEach(func() {
		kubectl.ValidateNoErrorsInLogs(CurrentGinkgoTestDescription().Duration)
		backgroundCancel()
	})

	AfterEach(func() {
		ExpectAllPodsTerminated(kubectl)
	})

	AfterAll(func() {
		kubectl.CloseSSHClient()
	})

	testHTTPRequest := func(clientPodLabel, url string) {
		pods, err := kubectl.GetPodNames(helpers.DefaultNamespace, clientPodLabel)
		ExpectWithOffset(1, err).Should(BeNil(), "cannot retrieve pod names by filter %q", testDSClient)
		// A DS with client is running in each node. So we try from each node
		// that can connect to the service.  To make sure that the cross-node
		// service connectivity is correct we tried 10 times, so balance in the
		// two nodes
		for _, pod := range pods {
			By("Making ten HTTP requests from %q to %q", pod, url)
			for i := 1; i <= 10; i++ {
				res := kubectl.ExecPodCmd(
					helpers.DefaultNamespace, pod,
					helpers.CurlFail(url))
				ExpectWithOffset(1, res).Should(helpers.CMDSuccess(),
					"Pod %q can not connect to service %q", pod, url)
			}
		}
	}

	waitPodsDs := func() {
		groups := []string{testDS, testDSClient}
		for _, pod := range groups {
			err := kubectl.WaitforPods(helpers.DefaultNamespace, fmt.Sprintf("-l %s", pod), helpers.HelperTimeout)
			ExpectWithOffset(1, err).Should(BeNil())
		}
	}

	Context("Checks ClusterIP Connectivity", func() {

		var (
			demoYAML    string
			echoSVCYAML string
		)

		BeforeAll(func() {

			demoYAML = helpers.ManifestGet(kubectl.BasePath(), "demo.yaml")
			echoSVCYAML = helpers.ManifestGet(kubectl.BasePath(), "echo-svc.yaml")
		})

		BeforeEach(func() {
			res := kubectl.ApplyDefault(demoYAML)
			res.ExpectSuccess("unable to apply %s", demoYAML)
			res = kubectl.ApplyDefault(echoSVCYAML)
			res.ExpectSuccess("unable to apply %s", echoSVCYAML)
		})

		AfterEach(func() {
			// Explicitly ignore result of deletion of resources to avoid incomplete
			// teardown if any step fails.
			_ = kubectl.Delete(demoYAML)
			_ = kubectl.Delete(echoSVCYAML)
		})

		It("Checks service on same node", func() {
			err := kubectl.WaitforPods(helpers.DefaultNamespace, "-l zgroup=testapp", helpers.HelperTimeout)
			Expect(err).Should(BeNil())
			clusterIP, _, err := kubectl.GetServiceHostPort(helpers.DefaultNamespace, serviceName)
			Expect(err).Should(BeNil(), "Cannot get service %s", serviceName)
			Expect(govalidator.IsIP(clusterIP)).Should(BeTrue(), "ClusterIP is not an IP")

			By("testing connectivity via cluster IP %s", clusterIP)
			monitorStop := kubectl.MonitorStart(helpers.KubeSystemNamespace, ciliumPodK8s1,
				"cluster-ip-same-node.log")
			status, err := kubectl.ExecInHostNetNS(context.TODO(), helpers.K8s1,
				helpers.CurlFail("http://%s/", clusterIP))
			monitorStop()
			Expect(err).To(BeNil(), "Cannot run curl in host netns")

			status.ExpectSuccess("cannot curl to service IP from host")
			ciliumPods, err := kubectl.GetCiliumPods(helpers.KubeSystemNamespace)
			Expect(err).To(BeNil(), "Cannot get cilium pods")
			for _, pod := range ciliumPods {
				service := kubectl.CiliumExec(pod, "cilium service list")
				service.ExpectSuccess("Cannot retrieve services on cilium Pod")
				service.ExpectContains(clusterIP, "ClusterIP is not present in the cilium service list")
			}
		}, 300)

		It("Checks service accessing itself (hairpin flow)", func() {
			err := kubectl.WaitforPods(helpers.DefaultNamespace, "-l name=echo", helpers.HelperTimeout)
			Expect(err).Should(BeNil())
			clusterIP, _, err := kubectl.GetServiceHostPort(helpers.DefaultNamespace, echoServiceName)
			Expect(err).Should(BeNil(), "Cannot get service %q ClusterIP", echoServiceName)
			Expect(govalidator.IsIP(clusterIP)).Should(BeTrue(), "ClusterIP is not an IP")

			url := fmt.Sprintf("http://%s/", clusterIP)
			testHTTPRequest(echoPodLabel, url)
		}, 300)
	})

	Context("Checks service across nodes", func() {

		var (
			demoYAML string
		)

		BeforeAll(func() {
			demoYAML = helpers.ManifestGet(kubectl.BasePath(), "demo_ds.yaml")
			res := kubectl.ApplyDefault(demoYAML)
			res.ExpectSuccess("Unable to apply %s", demoYAML)
		})

		AfterAll(func() {
			// Explicitly ignore result of deletion of resources to avoid incomplete
			// teardown if any step fails.
			_ = kubectl.Delete(demoYAML)
			ExpectAllPodsTerminated(kubectl)
		})

		It("Checks ClusterIP Connectivity", func() {
			waitPodsDs()
			service := "testds-service"

			clusterIP, _, err := kubectl.GetServiceHostPort(helpers.DefaultNamespace, service)
			Expect(err).Should(BeNil(), "Cannot get service %s", service)
			Expect(govalidator.IsIP(clusterIP)).Should(BeTrue(), "ClusterIP is not an IP")

			url := fmt.Sprintf("http://%s/", clusterIP)
			testHTTPRequest(testDSClient, url)
		})

		doRequests := func(url string, count int, fromPod string) {
			By("Making %d HTTP requests from %s to %q", count, fromPod, url)
			for i := 1; i <= count; i++ {
				res, err := kubectl.ExecInHostNetNS(context.TODO(), fromPod, helpers.CurlFail(url))
				ExpectWithOffset(1, err).To(BeNil(), "Cannot run curl in host netns")
				ExpectWithOffset(1, res).Should(helpers.CMDSuccess(),
					"%s host can not connect to service %q", fromPod, url)
			}
		}

		doRequestsFromOutsideClient := func(url string, count int) {
			ssh := helpers.GetVagrantSSHMeta(helpers.K8s1VMName())
			By("Making %d HTTP requests from outside cluster to %q", count, url)
			for i := 1; i <= count; i++ {
				res := ssh.ContainerExec("client-from-outside", helpers.CurlFail(url))
				ExpectWithOffset(1, res).Should(helpers.CMDSuccess(),
					"Can not connect to service %q from outside cluster", url)
			}
		}

		testNodePort := func(bpfNodePort bool) {
			var data v1.Service
			getURL := func(host string, port int32) string {
				return fmt.Sprintf("http://%s",
					net.JoinHostPort(host, fmt.Sprintf("%d", port)))
			}

			waitPodsDs()

			err := kubectl.Get(helpers.DefaultNamespace, "service test-nodeport").Unmarshal(&data)
			Expect(err).Should(BeNil(), "Can not retrieve service")
			url := getURL(data.Spec.ClusterIP, data.Spec.Ports[0].Port)
			testHTTPRequest(testDSClient, url)

			// From host via localhost IP
			// TODO: IPv6
			count := 10
			url = getURL("127.0.0.1", data.Spec.Ports[0].NodePort)
			doRequests(url, count, helpers.K8s1)

			url = getURL(helpers.K8s1Ip, data.Spec.Ports[0].NodePort)
			doRequests(url, count, helpers.K8s1)

			url = getURL(helpers.K8s2Ip, data.Spec.Ports[0].NodePort)
			doRequests(url, count, helpers.K8s1)

			// From pod via node IPs
			url = getURL(helpers.K8s1Ip, data.Spec.Ports[0].NodePort)
			testHTTPRequest(testDSClient, url)
			url = getURL(helpers.K8s2Ip, data.Spec.Ports[0].NodePort)
			testHTTPRequest(testDSClient, url)

			if bpfNodePort {
				// From host via local cilium_host
				localCiliumHostIPv4, err := kubectl.GetCiliumHostIPv4(context.TODO(), helpers.K8s1)
				Expect(err).Should(BeNil(), "Cannot retrieve local cilium_host ipv4")
				url = getURL(localCiliumHostIPv4, data.Spec.Ports[0].NodePort)
				doRequests(url, count, helpers.K8s1)

				// From host via remote cilium_host
				remoteCiliumHostIPv4, err := kubectl.GetCiliumHostIPv4(context.TODO(), helpers.K8s2)
				Expect(err).Should(BeNil(), "Cannot retrieve remote cilium_host ipv4")
				url = getURL(remoteCiliumHostIPv4, data.Spec.Ports[0].NodePort)
				doRequests(url, count, helpers.K8s1)

				// From pod via loopback (host reachable services)
				url = getURL("127.0.0.1", data.Spec.Ports[0].NodePort)
				testHTTPRequest(testDSClient, url)

				// From pod via local cilium_host
				url = getURL(localCiliumHostIPv4, data.Spec.Ports[0].NodePort)
				testHTTPRequest(testDSClient, url)

				// From pod via remote cilium_host
				url = getURL(remoteCiliumHostIPv4, data.Spec.Ports[0].NodePort)
				testHTTPRequest(testDSClient, url)
			}
		}

		It("Tests NodePort (kube-proxy)", func() {
			testNodePort(false)
		})

		Context("with L7 policy", func() {
			var (
				demoPolicy string
			)

			BeforeAll(func() {
				demoPolicy = helpers.ManifestGet(kubectl.BasePath(), "l7-policy-demo.yaml")
			})

			AfterAll(func() {
				// Explicitly ignore result of deletion of resources to avoid incomplete
				// teardown if any step fails.
				_ = kubectl.Delete(demoPolicy)
			})

			It("Tests NodePort with L7 Policy", func() {
				applyPolicy(demoPolicy)
				testNodePort(false)
			})
		})

		SkipContextIf(helpers.DoesNotRunOnNetNext, "Tests NodePort BPF", func() {
			// TODO(brb) Add with L7 policy test cases after GH#8971 has been fixed

			nativeDev := "enp0s8"

			BeforeAll(func() {
				enableBackgroundReport = false
			})

			AfterAll(func() {
				enableBackgroundReport = true
				// Remove NodePort programs (GH#8873)
				pods, err := kubectl.GetCiliumPods(helpers.KubeSystemNamespace)
				Expect(err).To(BeNil(), "Cannot retrieve Cilium pods")
				for _, pod := range pods {
					ret := kubectl.CiliumExec(pod, "tc filter del dev "+nativeDev+" ingress")
					Expect(ret.WasSuccessful()).Should(BeTrue(), "Cannot remove ingress bpf_netdev on %s", pod)
					ret = kubectl.CiliumExec(pod, "tc filter del dev "+nativeDev+" egress")
					Expect(ret.WasSuccessful()).Should(BeTrue(), "Cannot remove egress bpf_netdev on %s", pod)
				}
				deleteCiliumDS(kubectl)
				// Deploy Cilium as the next test expects it to be up and running
				DeployCiliumAndDNS(kubectl)
			})

			It("Tests with vxlan", func() {
				deleteCiliumDS(kubectl)

				DeployCiliumOptionsAndDNS(kubectl, []string{
					"--set global.nodePort.enabled=true",
					"--set global.nodePort.device=" + nativeDev,
				})

				testNodePort(true)
			})

			It("Tests with direct routing", func() {
				deleteCiliumDS(kubectl)
				DeployCiliumOptionsAndDNS(kubectl, []string{
					"--set global.nodePort.enabled=true",
					"--set global.nodePort.device=" + nativeDev,
					"--set global.tunnel=disabled",
					"--set global.autoDirectNodeRoutes=true",
				})

				testNodePort(true)
			})

			Context("Tests with MetalLB", func() {
				var (
					metalLB string
				)

				BeforeAll(func() {
					// Will allocate LoadBalancer IPs from 192.168.36.{240-250} range
					metalLB = helpers.ManifestGet(kubectl.BasePath(), "metallb.yaml")
					res := kubectl.ApplyDefault(metalLB)
					res.ExpectSuccess("Unable to apply %s", metalLB)
				})

				AfterAll(func() {
					_ = kubectl.Delete(metalLB)
				})

				It("Connectivity to endpoint via LB", func() {
					lbIP, err := kubectl.GetLoadBalancerIP(
						helpers.DefaultNamespace, "test-lb", 30*time.Second)
					Expect(err).Should(BeNil(), "Cannot retrieve loadbalancer IP for test-lb")

					doRequestsFromOutsideClient("http://"+lbIP, 10)
					doRequests("http://"+lbIP, 10, helpers.K8s1)
					doRequests("http://"+lbIP, 10, helpers.K8s2)
				})
			})
		})

	})

	//TODO: Check service with IPV6

	Context("External services", func() {

		var (
			expectedCIDR = "198.49.23.144/32"
			podName      = "toservices"

			endpointPath      string
			podPath           string
			policyPath        string
			policyLabeledPath string
			servicePath       string
		)

		BeforeAll(func() {
			endpointPath = helpers.ManifestGet(kubectl.BasePath(), "external_endpoint.yaml")
			podPath = helpers.ManifestGet(kubectl.BasePath(), "external_pod.yaml")
			policyPath = helpers.ManifestGet(kubectl.BasePath(), "external-policy.yaml")
			policyLabeledPath = helpers.ManifestGet(kubectl.BasePath(), "external-policy-labeled.yaml")
			servicePath = helpers.ManifestGet(kubectl.BasePath(), "external_service.yaml")

			kubectl.ApplyDefault(servicePath).ExpectSuccess("cannot install external service")
			kubectl.ApplyDefault(podPath).ExpectSuccess("cannot install pod path")

			err := kubectl.WaitforPods(helpers.DefaultNamespace, "", helpers.HelperTimeout)
			Expect(err).To(BeNil(), "Pods are not ready after timeout")

			err = kubectl.CiliumEndpointWaitReady()
			Expect(err).To(BeNil(), "Endpoints are not ready after timeout")
		})

		AfterAll(func() {
			_ = kubectl.Delete(servicePath)
			_ = kubectl.Delete(podPath)

			ExpectAllPodsTerminated(kubectl)
		})

		AfterEach(func() {
			_ = kubectl.Delete(policyLabeledPath)
			_ = kubectl.Delete(policyPath)
			_ = kubectl.Delete(endpointPath)
		})

		validateEgress := func() {
			By("Checking that toServices CIDR is plumbed into CEP")
			Eventually(func() string {
				res := kubectl.Exec(fmt.Sprintf(
					"%s -n %s get cep %s -o json",
					helpers.KubectlCmd,
					helpers.DefaultNamespace,
					podName))
				ExpectWithOffset(1, res).Should(helpers.CMDSuccess(), "cannot get Cilium endpoint")
				data, err := res.Filter(`{.status.policy.egress}`)
				ExpectWithOffset(1, err).To(BeNil(), "unable to get endpoint %s metadata", podName)
				return data.String()
			}, 2*time.Minute, 2*time.Second).Should(ContainSubstring(expectedCIDR))
		}

		validateEgressAfterDeletion := func() {
			By("Checking that toServices CIDR is no longer plumbed into CEP")
			Eventually(func() string {
				res := kubectl.Exec(fmt.Sprintf(
					"%s -n %s get cep %s -o json",
					helpers.KubectlCmd,
					helpers.DefaultNamespace,
					podName))
				ExpectWithOffset(1, res).Should(helpers.CMDSuccess(), "cannot get Cilium endpoint")
				data, err := res.Filter(`{.status.policy.egress}`)
				ExpectWithOffset(1, err).To(BeNil(), "unable to get endpoint %s metadata", podName)
				return data.String()
			}, 2*time.Minute, 2*time.Second).ShouldNot(ContainSubstring(expectedCIDR))
		}

		It("To Services first endpoint creation", func() {
			res := kubectl.ApplyDefault(endpointPath)
			res.ExpectSuccess()

			applyPolicy(policyPath)
			validateEgress()

			kubectl.Delete(policyPath)
			kubectl.Delete(endpointPath)
			validateEgressAfterDeletion()
		})

		It("To Services first policy", func() {
			applyPolicy(policyPath)
			res := kubectl.ApplyDefault(endpointPath)
			res.ExpectSuccess()

			validateEgress()

			kubectl.Delete(policyPath)
			kubectl.Delete(endpointPath)
			validateEgressAfterDeletion()
		})

		It("To Services first endpoint creation match service by labels", func() {
			By("Creating Kubernetes Endpoint")
			res := kubectl.ApplyDefault(endpointPath)
			res.ExpectSuccess()

			applyPolicy(policyLabeledPath)

			validateEgress()

			kubectl.Delete(policyLabeledPath)
			kubectl.Delete(endpointPath)
			validateEgressAfterDeletion()
		})

		It("To Services first policy, match service by labels", func() {
			applyPolicy(policyLabeledPath)

			By("Creating Kubernetes Endpoint")
			res := kubectl.ApplyDefault(endpointPath)
			res.ExpectSuccess()

			validateEgress()

			kubectl.Delete(policyLabeledPath)
			kubectl.Delete(endpointPath)
			validateEgressAfterDeletion()
		})
	})

	// FIXME: to test external IPs one needs to setup a routing to the VMs
	//        a manual test can be achieved by running
	//        Disable kube-proxy (not possible with the current CI)
	//        sudo service kube-proxy stop && sudo su -c 'iptables-save | grep -v KUBE | iptables-restore'
	//        start cilium with --enable-node-port --device=enp0s9
	//        kubectl label node k8s1 cilium.io/ci-node=k8s1
	//        kubectl apply -f test/k8sT/manifests/demo.yaml
	//        kubectl apply -f test/k8sT/manifests/external-ips-service.yaml
	//        (outside of the VM add a route for the external IP via the IP
	//        of the VM where the service is installed)
	//        $ ip r a 192.0.2.233 via 192.168.34.11
	//        $ curl 192.0.2.233:82
	//        <html><body><h1>It works!</h1></body></html>
	//
	//        Don't forget kubectl label node k8s1 cilium.io/ci-node=k8s1
	//        Expected behavior inside VMs with kube-proxy
	//        (make sure the external-ips-port is the same as the one created by node port)
	//        NOTES: When setting a external-ips-service with a port already allocated, for example 31388
	//        kube-proxy will not allow it and even print a warning in its logs.
	//        $ kubectl apply -f examples/misc/guestbook/
	//        $ kubectl apply -f test/k8sT/manifests/external-ips-service.yaml
	//
	//        Expected behavior inside VMs node with kube-proxy
	// 		  NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP                 PORT(S)        AGE
	// 	      external-ips-service   ClusterIP   172.20.0.224   192.0.2.233,192.168.34.11   31388/TCP      6s
	// 	      guestbook-nodeport     NodePort    172.20.0.171   <none>                      82:31388/TCP   6s
	//
	//        $ curl 172.20.0.224:31388 # works (shows `<html><body><h1>It works!</h1></body></html>`)
	//        $ curl 192.0.2.233:31388 # fails has expected because it should not work inside a VM
	//        $ (outside VM) curl 192.0.2.233:31388 # works (shows `<html><body><h1>It works!</h1></body></html>`)
	//        $ curl 192.168.34.11:31388 | grep '<title>Guestbook' # returns guestbook because kube-proxy set 31388
	//            for node-port service first (should have returned same as 192.0.2.233:31388)
	//        $ curl 127.0.0.1:31388 | grep '<title>Guestbook' # returns guestbook has expected because kube-proxy set 31388 for node-port service first
	//        $ curl 172.20.0.171:82 | grep '<title>Guestbook' # works as expected
	//
	// Context("External IPs services", func() {
	//
	// 	var (
	// 		externalIP                              = "192.0.2.233"
	// 		expectedCIDR                            = externalIP + "/32"
	// 		podName                                 = "toservices"
	// 		podPath, policyLabeledPath, servicePath string
	//
	// 		// shouldConnect asserts that srcPod can connect to dst.
	// 		shouldConnect = func(srcPod, dst string) {
	// 			By("Checking that %q can connect to %q", srcPod, dst)
	// 			res := kubectl.ExecPodCmd(helpers.DefaultNamespace, srcPod, fmt.Sprintf("sh -c 'rm -f index.html && wget %s'", dst))
	// 			res.ExpectSuccess("Unable to connect from %q to %q", srcPod, dst)
	// 		}
	// 	)
	//
	// 	BeforeAll(func() {
	// 		podPath = helpers.ManifestGet(kubectl.BasePath(), "demo.yaml")
	// 		policyLabeledPath = helpers.ManifestGet(kubectl.BasePath(), "external-policy-labeled.yaml")
	// 		servicePath = helpers.ManifestGet(kubectl.BasePath(), "external-ips-service.yaml")
	//
	// 		localExec := helpers.CreateLocalExecutor(os.Environ())
	// 		localExec.Exec(fmt.Sprintf("sudo ip route add %s via %s", externalIP, helpers.K8s1Ip))
	// 		kubectl.Apply(servicePath).ExpectSuccess("cannot install external service")
	// 		kubectl.Apply(podPath).ExpectSuccess("cannot install pod path")
	//
	// 		err := kubectl.WaitforPods(helpers.DefaultNamespace, "", helpers.HelperTimeout)
	// 		Expect(err).To(BeNil(), "Pods are not ready after timeout")
	//
	// 		err = kubectl.CiliumEndpointWaitReady()
	// 		Expect(err).To(BeNil(), "Endpoints are not ready after timeout")
	// 	})
	//
	// 	AfterAll(func() {
	// 		_ = kubectl.Delete(servicePath)
	// 		_ = kubectl.Delete(podPath)
	//
	// 		ExpectAllPodsTerminated(kubectl)
	// 		localExec := helpers.CreateLocalExecutor(os.Environ())
	// 		localExec.Exec(fmt.Sprintf("sudo ip route delete %s via %s", externalIP, helpers.K8s1Ip))
	// 	})
	//
	// 	AfterEach(func() {
	// 		_ = kubectl.Delete(policyLabeledPath)
	// 	})
	//
	// 	validateEgress := func() {
	// 		By("Checking that toServices CIDR is plumbed into CEP")
	// 		Eventually(func() string {
	// 			res := kubectl.Exec(fmt.Sprintf(
	// 				"%s -n %s get cep %s -o json",
	// 				helpers.KubectlCmd,
	// 				helpers.DefaultNamespace,
	// 				podName))
	// 			ExpectWithOffset(1, res).Should(helpers.CMDSuccess(), "cannot get Cilium endpoint")
	// 			data, err := res.Filter(`{.status.policy.egress}`)
	// 			ExpectWithOffset(1, err).To(BeNil(), "unable to get endpoint %s metadata", podName)
	// 			return data.String()
	// 		}, 2*time.Minute, 2*time.Second).Should(ContainSubstring(expectedCIDR))
	// 	}
	//
	// 	validateEgressAfterDeletion := func() {
	// 		By("Checking that toServices CIDR is no longer plumbed into CEP")
	// 		Eventually(func() string {
	// 			res := kubectl.Exec(fmt.Sprintf(
	// 				"%s -n %s get cep %s -o json",
	// 				helpers.KubectlCmd,
	// 				helpers.DefaultNamespace,
	// 				podName))
	// 			ExpectWithOffset(1, res).Should(helpers.CMDSuccess(), "cannot get Cilium endpoint")
	// 			data, err := res.Filter(`{.status.policy.egress}`)
	// 			ExpectWithOffset(1, err).To(BeNil(), "unable to get endpoint %s metadata", podName)
	// 			return data.String()
	// 		}, 2*time.Minute, 2*time.Second).ShouldNot(ContainSubstring(expectedCIDR))
	// 	}
	//
	// 	It("Connects to external IPs", func() {
	// 		shouldConnect(podName, externalIP)
	// 	})
	//
	// 	It("Connects to service IP backed by external IPs", func() {
	// 		err := kubectl.WaitForKubeDNSEntry("external-ips-service", helpers.DefaultNamespace)
	// 		Expect(err).To(BeNil(), "DNS entry is not ready after timeout")
	// 		shouldConnect(podName, serviceName)
	// 	})
	//
	// 	It("To Services first endpoint creation match service by labels", func() {
	// 		By("Creating Kubernetes Endpoint")
	// 		applyPolicy(policyLabeledPath)
	//
	// 		validateEgress()
	//
	// 		kubectl.Delete(policyLabeledPath)
	// 		validateEgressAfterDeletion()
	// 	})
	//
	// })

	Context("Bookinfo Demo", func() {

		var (
			bookinfoV1YAML, bookinfoV2YAML string
			resourceYAMLs                  []string
			policyPath                     string
		)

		BeforeEach(func() {

			bookinfoV1YAML = helpers.ManifestGet(kubectl.BasePath(), "bookinfo-v1.yaml")
			bookinfoV2YAML = helpers.ManifestGet(kubectl.BasePath(), "bookinfo-v2.yaml")
			policyPath = helpers.ManifestGet(kubectl.BasePath(), "cnp-specs.yaml")

			resourceYAMLs = []string{bookinfoV1YAML, bookinfoV2YAML}

			for _, resourcePath := range resourceYAMLs {
				By("Creating objects in file %q", resourcePath)
				res := kubectl.Create(resourcePath)
				res.ExpectSuccess("unable to create resource %q", resourcePath)
			}
		})

		AfterEach(func() {

			// Explicitly do not check result to avoid having assertions in AfterEach.
			_ = kubectl.Delete(policyPath)

			for _, resourcePath := range resourceYAMLs {
				By("Deleting resource %s", resourcePath)
				// Explicitly do not check result to avoid having assertions in AfterEach.
				_ = kubectl.Delete(resourcePath)
			}
		})

		It("Tests bookinfo demo", func() {

			// We use wget in this test because the Istio apps do not provide curl.
			wgetCommand := fmt.Sprintf("wget --tries=2 --connect-timeout %d", helpers.CurlConnectTimeout)

			version := "version"
			v1 := "v1"

			productPage := "productpage"
			reviews := "reviews"
			ratings := "ratings"
			details := "details"
			dnsChecks := []string{productPage, reviews, ratings, details}
			app := "app"
			health := "health"
			ratingsPath := "ratings/0"

			apiPort := "9080"

			podNameFilter := "{.items[*].metadata.name}"

			// shouldConnect asserts that srcPod can connect to dst.
			shouldConnect := func(srcPod, dst string) {
				By("Checking that %q can connect to %q", srcPod, dst)
				res := kubectl.ExecPodCmd(
					helpers.DefaultNamespace, srcPod, fmt.Sprintf("%s %s", wgetCommand, dst))
				res.ExpectSuccess("Unable to connect from %q to %q", srcPod, dst)
			}

			// shouldNotConnect asserts that srcPod cannot connect to dst.
			shouldNotConnect := func(srcPod, dst string) {
				By("Checking that %q cannot connect to %q", srcPod, dst)
				res := kubectl.ExecPodCmd(
					helpers.DefaultNamespace, srcPod, fmt.Sprintf("%s %s", wgetCommand, dst))
				res.ExpectFail("Was able to connect from %q to %q, but expected no connection: %s", srcPod, dst, res.CombineOutput())
			}

			// formatLabelArgument formats the provided key-value pairs as labels for use in
			// querying Kubernetes.
			formatLabelArgument := func(firstKey, firstValue string, nextLabels ...string) string {
				baseString := fmt.Sprintf("-l %s=%s", firstKey, firstValue)
				if nextLabels == nil {
					return baseString
				} else if len(nextLabels)%2 != 0 {
					Fail("must provide even number of arguments for label key-value pairings")
				} else {
					for i := 0; i < len(nextLabels); i += 2 {
						baseString = fmt.Sprintf("%s,%s=%s", baseString, nextLabels[i], nextLabels[i+1])
					}
				}
				return baseString
			}

			// formatAPI is a helper function which formats a URI to access.
			formatAPI := func(service, port, resource string) string {
				target := fmt.Sprintf(
					"%s.%s.svc.cluster.local:%s",
					service, helpers.DefaultNamespace, port)
				if resource != "" {
					return fmt.Sprintf("%s/%s", target, resource)
				}
				return target
			}

			By("Waiting for pods to be ready")
			err := kubectl.WaitforPods(helpers.DefaultNamespace, "-l zgroup=bookinfo", helpers.HelperTimeout)
			Expect(err).Should(BeNil(), "Pods are not ready after timeout")

			err = kubectl.CiliumEndpointWaitReady()
			ExpectWithOffset(1, err).To(BeNil(), "Endpoints are not ready after timeout")

			By("Waiting for services to be ready")
			for _, service := range []string{details, ratings, reviews, productPage} {
				err = kubectl.WaitForServiceEndpoints(
					helpers.DefaultNamespace, "", service,
					helpers.HelperTimeout)
				Expect(err).Should(BeNil(), "Service %q is not ready after timeout", service)
			}
			By("Validating DNS without Policy")
			for _, name := range dnsChecks {
				err = kubectl.WaitForKubeDNSEntry(name, helpers.DefaultNamespace)
				Expect(err).To(BeNil(), "DNS entry is not ready after timeout")
			}

			By("All pods should be able to connect without policy")

			reviewsPodV1, err := kubectl.GetPods(helpers.DefaultNamespace, formatLabelArgument(app, reviews, version, v1)).Filter(podNameFilter)
			Expect(err).Should(BeNil(), "cannot get reviewsV1 pods")
			productpagePodV1, err := kubectl.GetPods(helpers.DefaultNamespace, formatLabelArgument(app, productPage, version, v1)).Filter(podNameFilter)
			Expect(err).Should(BeNil(), "cannot get productpageV1 pods")

			shouldConnect(reviewsPodV1.String(), formatAPI(ratings, apiPort, health))
			shouldConnect(reviewsPodV1.String(), formatAPI(ratings, apiPort, ratingsPath))

			shouldConnect(productpagePodV1.String(), formatAPI(details, apiPort, health))
			shouldConnect(productpagePodV1.String(), formatAPI(ratings, apiPort, health))
			shouldConnect(productpagePodV1.String(), formatAPI(ratings, apiPort, ratingsPath))

			policyCmd := "cilium policy get io.cilium.k8s.policy.name=cnp-specs"

			By("Importing policy")

			_, err = kubectl.CiliumPolicyAction(helpers.DefaultNamespace, policyPath, helpers.KubectlCreate, helpers.HelperTimeout)
			Expect(err).Should(BeNil(), "Error creating policy %q", policyPath)

			By("Checking that policies were correctly imported into Cilium")

			ciliumPodK8s1, err = kubectl.GetCiliumPodOnNodeWithLabel(helpers.KubeSystemNamespace, helpers.K8s1)
			Expect(err).Should(BeNil(), "Cannot get cilium pod on k8s1")
			res := kubectl.ExecPodCmd(helpers.KubeSystemNamespace, ciliumPodK8s1, policyCmd)
			res.ExpectSuccess("Policy %s is not imported", policyCmd)

			By("Validating DNS with Policy loaded")
			for _, name := range dnsChecks {
				err = kubectl.WaitForKubeDNSEntry(name, helpers.DefaultNamespace)
				Expect(err).To(BeNil(), "DNS entry is not ready after timeout")
			}

			By("After policy import")
			shouldConnect(reviewsPodV1.String(), formatAPI(ratings, apiPort, health))
			shouldNotConnect(reviewsPodV1.String(), formatAPI(ratings, apiPort, ratingsPath))

			shouldConnect(productpagePodV1.String(), formatAPI(details, apiPort, health))

			shouldNotConnect(productpagePodV1.String(), formatAPI(ratings, apiPort, health))
			shouldNotConnect(productpagePodV1.String(), formatAPI(ratings, apiPort, ratingsPath))
		})
	})
})
