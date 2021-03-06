package uri

import (
	"artion-api-graphql/internal/types"
	ipfsapi "github.com/ipfs/go-ipfs-api"
	"github.com/onsi/gomega"
	"testing"
)

func TestDataUri(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	downloader := Downloader{
		ipfsShell: nil,
	}
	data,err := downloader.GetJsonMetadata("data:application/json;base64,eyJuYW1lIjogIml0ZW0gIzkyOTIiLCAiZGVzY3JpcHRpb24iOiAiUmFyaXR5IHRpZXIgMSwgbm9uIG1hZ2ljYWwsIGl0ZW0gY3JhZnRpbmcuIiwgImltYWdlIjogImRhdGE6aW1hZ2Uvc3ZnK3htbDtiYXNlNjQsUEhOMlp5QjRiV3h1Y3owaWFIUjBjRG92TDNkM2R5NTNNeTV2Y21jdk1qQXdNQzl6ZG1jaUlIQnlaWE5sY25abFFYTndaV04wVW1GMGFXODlJbmhOYVc1WlRXbHVJRzFsWlhRaUlIWnBaWGRDYjNnOUlqQWdNQ0F6TlRBZ016VXdJajQ4YzNSNWJHVStMbUpoYzJVZ2V5Qm1hV3hzT2lCM2FHbDBaVHNnWm05dWRDMW1ZVzFwYkhrNklITmxjbWxtT3lCbWIyNTBMWE5wZW1VNklERTBjSGc3SUgwOEwzTjBlV3hsUGp4eVpXTjBJSGRwWkhSb1BTSXhNREFsSWlCb1pXbG5hSFE5SWpFd01DVWlJR1pwYkd3OUltSnNZV05ySWlBdlBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeU1DSWdZMnhoYzNNOUltSmhjMlVpUG1OaGRHVm5iM0o1SUVGeWJXOXlQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJME1DSWdZMnhoYzNNOUltSmhjMlVpUG01aGJXVWdRbkpsWVhOMGNHeGhkR1U4TDNSbGVIUStQSFJsZUhRZ2VEMGlNVEFpSUhrOUlqWXdJaUJqYkdGemN6MGlZbUZ6WlNJK1kyOXpkQ0F5TURCbmNEd3ZkR1Y0ZEQ0OGRHVjRkQ0I0UFNJeE1DSWdlVDBpT0RBaUlHTnNZWE56UFNKaVlYTmxJajUzWldsbmFIUWdNekJzWWp3dmRHVjRkRDQ4ZEdWNGRDQjRQU0l4TUNJZ2VUMGlNVEF3SWlCamJHRnpjejBpWW1GelpTSStjSEp2Wm1samFXVnVZM2tnVFdWa2FYVnRQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeE1qQWlJR05zWVhOelBTSmlZWE5sSWo1aGNtMXZjaUJpYjI1MWN5QTFQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeE5EQWlJR05zWVhOelBTSmlZWE5sSWo1dFlYZ2daR1Y0SURNOEwzUmxlSFErUEhSbGVIUWdlRDBpTVRBaUlIazlJakUyTUNJZ1kyeGhjM005SW1KaGMyVWlQbkJsYm1Gc2RIa2dMVFE4TDNSbGVIUStQSFJsZUhRZ2VEMGlNVEFpSUhrOUlqRTRNQ0lnWTJ4aGMzTTlJbUpoYzJVaVBuTndaV3hzSUdaaGFXeDFjbVVnTWpVbFBDOTBaWGgwUGp4MFpYaDBJSGc5SWpFd0lpQjVQU0l5TURBaUlHTnNZWE56UFNKaVlYTmxJajVrWlhOamNtbHdkR2x2YmlCSmRDQmpiMjFsY3lCM2FYUm9JR0VnYUdWc2JXVjBJR0Z1WkNCbmNtVmhkbVZ6TGp3dmRHVjRkRDQ4ZEdWNGRDQjRQU0l4TUNJZ2VUMGlNakl3SWlCamJHRnpjejBpWW1GelpTSStZM0poWm5SbFpDQmllU0F4TXpZeU1qZ3dQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeU5EQWlJR05zWVhOelBTSmlZWE5sSWo1amNtRm1kR1ZrSUdGMElERTJNekkxTnpRMk1UZzhMM1JsZUhRK1BDOXpkbWMrIn0=")
	g.Expect(err).To(gomega.BeNil())
	g.Expect(data.Name).To(gomega.Equal("item #9292"))
	g.Expect(*data.Description).To(gomega.Equal("Rarity tier 1, non magical, item crafting."))
}

func TestHttp(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	downloader := Downloader{
		ipfsShell: nil,
	}
	data,err := downloader.GetJsonMetadata("https://gist.githubusercontent.com/hkalina/0fdb14b893ad912be79fc6ae0ec5d86b/raw/50ad50fa7c9292bcda1a6622431765ef0e874f32/test.json")
	g.Expect(err).To(gomega.BeNil())
	g.Expect(data.Name).To(gomega.Equal("item #9292"))
	g.Expect(*data.Description).To(gomega.Equal("Rarity tier 1, non magical, item crafting."))
}

func TestIpfs(t *testing.T) { // requires locally running IPFS node
	g := gomega.NewGomegaWithT(t)

	downloader := Downloader{
		ipfsShell: ipfsapi.NewShell("localhost:5001"),
	}
	data,err := downloader.GetJsonMetadata("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi")
	g.Expect(err).To(gomega.BeNil())
	g.Expect(data.Name).To(gomega.Equal("item #9292"))
	g.Expect(*data.Description).To(gomega.Equal("Rarity tier 1, non magical, item crafting."))
}

func TestDataUriMimetype(t *testing.T) { // requires locally running IPFS node
	g := gomega.NewGomegaWithT(t)

	downloader := Downloader{
		ipfsShell: nil,
	}
	data,err := downloader.GetImage("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHByZXNlcnZlQXNwZWN0UmF0aW89InhNaW5ZTWluIG1lZXQiIHZpZXdCb3g9IjAgMCAzNTAgMzUwIj48c3R5bGU+LmJhc2UgeyBmaWxsOiB3aGl0ZTsgZm9udC1mYW1pbHk6IHNlcmlmOyBmb250LXNpemU6IDE0cHg7IH08L3N0eWxlPjxyZWN0IHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIGZpbGw9ImJsYWNrIiAvPjx0ZXh0IHg9IjEwIiB5PSIyMCIgY2xhc3M9ImJhc2UiPmNhdGVnb3J5IEFybW9yPC90ZXh0Pjx0ZXh0IHg9IjEwIiB5PSI0MCIgY2xhc3M9ImJhc2UiPm5hbWUgQnJlYXN0cGxhdGU8L3RleHQ+PHRleHQgeD0iMTAiIHk9IjYwIiBjbGFzcz0iYmFzZSI+Y29zdCAyMDBncDwvdGV4dD48dGV4dCB4PSIxMCIgeT0iODAiIGNsYXNzPSJiYXNlIj53ZWlnaHQgMzBsYjwvdGV4dD48dGV4dCB4PSIxMCIgeT0iMTAwIiBjbGFzcz0iYmFzZSI+cHJvZmljaWVuY3kgTWVkaXVtPC90ZXh0Pjx0ZXh0IHg9IjEwIiB5PSIxMjAiIGNsYXNzPSJiYXNlIj5hcm1vciBib251cyA1PC90ZXh0Pjx0ZXh0IHg9IjEwIiB5PSIxNDAiIGNsYXNzPSJiYXNlIj5tYXggZGV4IDM8L3RleHQ+PHRleHQgeD0iMTAiIHk9IjE2MCIgY2xhc3M9ImJhc2UiPnBlbmFsdHkgLTQ8L3RleHQ+PHRleHQgeD0iMTAiIHk9IjE4MCIgY2xhc3M9ImJhc2UiPnNwZWxsIGZhaWx1cmUgMjUlPC90ZXh0Pjx0ZXh0IHg9IjEwIiB5PSIyMDAiIGNsYXNzPSJiYXNlIj5kZXNjcmlwdGlvbiBJdCBjb21lcyB3aXRoIGEgaGVsbWV0IGFuZCBncmVhdmVzLjwvdGV4dD48dGV4dCB4PSIxMCIgeT0iMjIwIiBjbGFzcz0iYmFzZSI+Y3JhZnRlZCBieSAxMzYyMjgwPC90ZXh0Pjx0ZXh0IHg9IjEwIiB5PSIyNDAiIGNsYXNzPSJiYXNlIj5jcmFmdGVkIGF0IDE2MzI1NzQ2MTg8L3RleHQ+PC9zdmc+")
	g.Expect(err).To(gomega.BeNil())
	g.Expect(data.Type).To(gomega.Equal(types.ImageTypeSvg))
}

func TestIpfsUriConversion(t *testing.T) { // requires locally running IPFS node
	g := gomega.NewGomegaWithT(t)
	d := Downloader{
		skipHttpGateways: true,
	}

	uri := d.getIpfsUri("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi"))

	uri = d.getIpfsUri("ipfs://QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi"))

	uri = d.getIpfsUri("https://gateway.pinata.cloud/ipfs/QmdnUo3E1B27aLLTn3LsNZZqWtRfKMNgHAYzwVE5iYrP9P")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmdnUo3E1B27aLLTn3LsNZZqWtRfKMNgHAYzwVE5iYrP9P"))

	uri = d.getIpfsUri("https://ipfs.io/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi?filename=TestToken.json")
	g.Expect(uri).To(gomega.Equal("/ipfs/QmTetVgMNVGj88s9NQuANyVmjMtZqhZDp8T21huiVGbfAi?filename=TestToken.json"))

	uri = d.getIpfsUri("https://artion6.mypinata.cloud/ipfs/Qma6nGVgxmf95FUP8zJPQo7vEzwuRhdC9i7QQAzFDFCuvx")
	g.Expect(uri).To(gomega.Equal("/ipfs/Qma6nGVgxmf95FUP8zJPQo7vEzwuRhdC9i7QQAzFDFCuvx"))

	uri = d.getIpfsUri("https://example.org/test.json")
	g.Expect(uri).To(gomega.Equal(""))
}