package uri

import (
	ipfsapi "github.com/ipfs/go-ipfs-api"
	"github.com/onsi/gomega"
	"testing"
)

func TestDataUri(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	downloader := Downloader{
		ipfsShell: nil,
	}
	data, _, err := downloader.GetFromDataUri("data:application/json;base64,eyJuYW1lIjogIml0ZW0gIzkyOTIiLCAiZGVzY3JpcHRpb24iOiAiUmFyaXR5IHRpZXIgMSwgbm9uIG1hZ2ljYWwsIGl0ZW0gY3JhZnRpbmcuIiwgImltYWdlIjogImRhdGE6aW1hZ2Uvc3ZnK3htbDtiYXNlNjQsUEhOMlp5QjRiV3h1Y3owaWFIUjBjRG92TDNkM2R5NTNNeTV2Y21jdk1qQXdNQzl6ZG1jaUlIQnlaWE5sY25abFFYTndaV04wVW1GMGFXODlJbmhOYVc1WlRXbHVJRzFsWlhRaUlIWnBaWGRDYjNnOUlqQWdNQ0F6TlRBZ016VXdJajQ4YzNSNWJHVStMbUpoYzJVZ2V5Qm1hV3hzT2lCM2FHbDBaVHNnWm05dWRDMW1ZVzFwYkhrNklITmxjbWxtT3lCbWIyNTBMWE5wZW1VNklERTBjSGc3SUgwOEwzTjBlV3hsUGp4eVpXTjBJSGRwWkhSb1BTSXhNREFsSWlCb1pXbG5hSFE5SWpFd01DVWlJR1pwYkd3OUltSnNZV05ySWlBdlBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeU1DSWdZMnhoYzNNOUltSmhjMlVpUG1OaGRHVm5iM0o1SUVGeWJXOXlQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJME1DSWdZMnhoYzNNOUltSmhjMlVpUG01aGJXVWdRbkpsWVhOMGNHeGhkR1U4TDNSbGVIUStQSFJsZUhRZ2VEMGlNVEFpSUhrOUlqWXdJaUJqYkdGemN6MGlZbUZ6WlNJK1kyOXpkQ0F5TURCbmNEd3ZkR1Y0ZEQ0OGRHVjRkQ0I0UFNJeE1DSWdlVDBpT0RBaUlHTnNZWE56UFNKaVlYTmxJajUzWldsbmFIUWdNekJzWWp3dmRHVjRkRDQ4ZEdWNGRDQjRQU0l4TUNJZ2VUMGlNVEF3SWlCamJHRnpjejBpWW1GelpTSStjSEp2Wm1samFXVnVZM2tnVFdWa2FYVnRQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeE1qQWlJR05zWVhOelBTSmlZWE5sSWo1aGNtMXZjaUJpYjI1MWN5QTFQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeE5EQWlJR05zWVhOelBTSmlZWE5sSWo1dFlYZ2daR1Y0SURNOEwzUmxlSFErUEhSbGVIUWdlRDBpTVRBaUlIazlJakUyTUNJZ1kyeGhjM005SW1KaGMyVWlQbkJsYm1Gc2RIa2dMVFE4TDNSbGVIUStQSFJsZUhRZ2VEMGlNVEFpSUhrOUlqRTRNQ0lnWTJ4aGMzTTlJbUpoYzJVaVBuTndaV3hzSUdaaGFXeDFjbVVnTWpVbFBDOTBaWGgwUGp4MFpYaDBJSGc5SWpFd0lpQjVQU0l5TURBaUlHTnNZWE56UFNKaVlYTmxJajVrWlhOamNtbHdkR2x2YmlCSmRDQmpiMjFsY3lCM2FYUm9JR0VnYUdWc2JXVjBJR0Z1WkNCbmNtVmhkbVZ6TGp3dmRHVjRkRDQ4ZEdWNGRDQjRQU0l4TUNJZ2VUMGlNakl3SWlCamJHRnpjejBpWW1GelpTSStZM0poWm5SbFpDQmllU0F4TXpZeU1qZ3dQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeU5EQWlJR05zWVhOelBTSmlZWE5sSWo1amNtRm1kR1ZrSUdGMElERTJNekkxTnpRMk1UZzhMM1JsZUhRK1BDOXpkbWMrIn0=")
	g.Expect(err).To(gomega.BeNil())
	g.Expect(data).To(gomega.ContainSubstring("item #9292"))
}

func TestHttp(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	downloader := Downloader{
		ipfsShell: nil,
	}
	data, _, err := downloader.GetFromHttp("https://0ni.mypinata.cloud/ipfs/QmRZBp1WRqhTnojxHHBUbtru8MnSq4EGsy3Qzwfhvn7mh9/567")
	g.Expect(err).To(gomega.BeNil())
	g.Expect(data).To(gomega.ContainSubstring("0ni  #567"))
}

func TestIpfs(t *testing.T) { // requires locally running IPFS node
	g := gomega.NewGomegaWithT(t)

	downloader := Downloader{
		ipfsShell: ipfsapi.NewShell("localhost:5001"),
	}
	data, _, err := downloader.GetFromIpfs("/ipfs/QmNuFBtzMEo3fBobV9CHHTVbAPYfNimuhfe91P22kcouBq")
	g.Expect(err).To(gomega.BeNil())
	g.Expect(data).To(gomega.ContainSubstring("Amaterasu"))
}

func TestIpfsUriConversion(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	d := Downloader{
		skipHttpGateways: true,
	}

	uri := d.GetIpfsUri("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi"))

	uri = d.GetIpfsUri("ipfs://QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi"))

	uri = d.GetIpfsUri("https://gateway.pinata.cloud/ipfs/QmdnUo3E1B27aLLTn3LsNZZqWtRfKMNgHAYzwVE5iYrP9P")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmdnUo3E1B27aLLTn3LsNZZqWtRfKMNgHAYzwVE5iYrP9P"))

	uri = d.GetIpfsUri("https://ipfs.io/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi?filename=TestToken.json")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi?filename=TestToken.json"))

	uri = d.GetIpfsUri("https://artion6.mypinata.cloud/ipfs/Qma6nGVgxmf95FUP8zJPQo7vEzwuRhdC9i7QQAzFDFCuvx")
	g.Expect(uri).To(gomega.Equal("/ipfs/Qma6nGVgxmf95FUP8zJPQo7vEzwuRhdC9i7QQAzFDFCuvx"))

	uri = d.GetIpfsUri("https://example.org/test.json")
	g.Expect(uri).To(gomega.Equal(""))
}