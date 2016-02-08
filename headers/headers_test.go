package headers_test

import (
	. "github.com/Pivotal-Japan/service-test/headers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var routeService *RouteServiceHeaders

var _ = Describe("Headers", func() {
	Describe("Setup properly", func() {
		Context("with valid header", func() {
			headers := &http.Header{}
			headers.Add("X-CF-Proxy-Signature", "ProxySign")
			headers.Add("X-CF-Forwarded-Url", "https://localhost.com")
			headers.Add("X-CF-Proxy-Metadata", "Proxy_Metada")

			routeService = NewRouteServiceHeaders()
			err := routeService.ParseHeaders(headers)
			It("Should be valid", func() {
				Expect(routeService.IsValidRequest()).To(Equal(true))
				Expect(err).To(BeNil())
			})
			It("Shoud return valid URL", func() {
				Expect(routeService.ParsedUrl.Host).To(Equal("localhost.com"))
			})

		})

	})
})
