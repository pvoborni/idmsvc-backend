package presenter

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hmsidm/internal/api/public"
	"github.com/hmsidm/internal/domain/model"
	"github.com/hmsidm/internal/interface/presenter"
	"github.com/lib/pq"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewTodoPresenter(t *testing.T) {
	assert.NotPanics(t, func() {
		NewDomainPresenter()
	})
}

type mynewerror struct{}

func (e *mynewerror) Error() string {
	return "mynewerror"
}

func TestGet(t *testing.T) {
	testUuid := uuid.New()
	type TestCaseGiven struct {
		Input  *model.Domain
		Output *public.ReadDomainResponse
	}
	type TestCaseExpected struct {
		Err    error
		Output *public.ReadDomainResponse
	}
	type TestCase struct {
		Name     string
		Given    TestCaseGiven
		Expected TestCaseExpected
	}
	testCases := []TestCase{
		{
			Name: "error when 'in' is nil",
			Given: TestCaseGiven{
				Input: nil,
			},
			Expected: TestCaseExpected{
				Err:    fmt.Errorf("'domain' cannot be nil"),
				Output: nil,
			},
		},
		{
			Name: "Success case",
			Given: TestCaseGiven{
				Input: &model.Domain{
					Model:                 gorm.Model{ID: 1},
					OrgId:                 "12345",
					DomainUuid:            testUuid,
					DomainName:            pointy.String("domain.example"),
					Type:                  pointy.Uint(model.DomainTypeIpa),
					AutoEnrollmentEnabled: pointy.Bool(true),
					IpaDomain: &model.Ipa{
						RealmName:    pointy.String("DOMAIN.EXAMPLE"),
						CaCerts:      []model.IpaCert{},
						Servers:      []model.IpaServer{},
						RealmDomains: pq.StringArray{"domain.example"},
					},
				},
			},
			Expected: TestCaseExpected{
				Err: nil,
				Output: &public.ReadDomainResponse{
					AutoEnrollmentEnabled: true,
					DomainUuid:            testUuid.String(),
					DomainName:            "domain.example",
					Type:                  public.DomainResponseType(model.DomainTypeString(model.DomainTypeIpa)),
					Ipa: public.DomainResponseIpa{
						RealmName:    "DOMAIN.EXAMPLE",
						CaCerts:      []public.DomainIpaCert{},
						Servers:      []public.DomainIpaServer{},
						RealmDomains: []string{"domain.example"},
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Log(testCase.Name)
		obj := NewDomainPresenter()
		output, err := obj.Get(testCase.Given.Input)
		if testCase.Expected.Err != nil {
			require.Error(t, err)
			assert.Equal(t, testCase.Expected.Err.Error(), err.Error())
			assert.Nil(t, output)
		} else {
			assert.NoError(t, err)
			assert.Equal(t,
				testCase.Expected.Output.DomainUuid,
				output.DomainUuid)
			assert.Equal(t,
				testCase.Expected.Output.DomainName,
				output.DomainName)
			assert.Equal(t,
				testCase.Expected.Output.Type,
				output.Type)
			assert.Equal(t,
				testCase.Expected.Output.AutoEnrollmentEnabled,
				output.AutoEnrollmentEnabled)
			assert.Equal(t,
				testCase.Expected.Output.Ipa.RealmName,
				output.Ipa.RealmName)
			assert.Equal(t,
				testCase.Expected.Output.Ipa.CaCerts,
				output.Ipa.CaCerts)
			assert.Equal(t,
				testCase.Expected.Output.Ipa.Servers,
				output.Ipa.Servers)
		}
	}
}

func TestCreate(t *testing.T) {
	type TestCaseExpected struct {
		Response *public.CreateDomainResponse
		Err      error
	}
	type TestCase struct {
		Name     string
		Given    *model.Domain
		Expected TestCaseExpected
	}
	var (
		obj presenter.DomainPresenter
	)
	obj = NewDomainPresenter()

	testCases := []TestCase{
		{
			Name:  "domain is nil",
			Given: nil,
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("domain cannot be nil"),
			},
		},
		{
			Name: "AutoenrollmentEnabled is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("AutoenrollmentEnabled cannot be nil"),
			},
		},
		{
			Name: "DomainName is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("DomainName cannot be nil"),
			},
		},
		{
			Name: "DomainType is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				Type:                  nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("DomainType cannot be nil"),
			},
		},
		{
			Name: "DomainType is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				Type:                  nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("DomainType cannot be nil"),
			},
		},
		{
			Name: "IpaDomain is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				Type:                  pointy.Uint(model.DomainTypeIpa),
				IpaDomain:             nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("IpaDomain cannot be nil"),
			},
		},
		{
			Name: "RealmName is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				Type:                  pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					RealmName: nil,
				},
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("RealmName cannot be nil"),
			},
		},
		{
			Name: "CaCerts is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				Type:                  pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					RealmName: pointy.String("DOMAIN.EXAMPLE"),
					CaCerts:   nil,
				},
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("CaCerts cannot be nil"),
			},
		},
		{
			Name: "Servers is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				Type:                  pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					RealmName: pointy.String("DOMAIN.EXAMPLE"),
					CaCerts:   []model.IpaCert{},
					Servers:   nil,
				},
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("Servers cannot be nil"),
			},
		},
		{
			Name: "Success scenario",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				Type:                  pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					RealmName: pointy.String("DOMAIN.EXAMPLE"),
					CaCerts:   []model.IpaCert{},
					Servers:   []model.IpaServer{},
				},
			},
			Expected: TestCaseExpected{
				Response: &public.DomainResponse{
					DomainName:            "domain.example",
					AutoEnrollmentEnabled: true,
					Type:                  model.DomainTypeIpaString,
					DomainUuid:            "00000000-0000-0000-0000-000000000000",
					Ipa: public.DomainResponseIpa{
						RealmName: "DOMAIN.EXAMPLE",
						CaCerts:   []public.DomainIpaCert{},
						Servers:   []public.DomainIpaServer{},
					},
				},
				Err: nil,
			},
		},
	}

	for _, testCase := range testCases {
		response, err := obj.Create(testCase.Given)
		if testCase.Expected.Err != nil {
			require.Error(t, err)
			assert.EqualError(t, err, testCase.Expected.Err.Error())
			assert.Nil(t, response)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.Expected.Response.DomainName, response.DomainName)
			assert.Equal(t, testCase.Expected.Response.AutoEnrollmentEnabled, response.AutoEnrollmentEnabled)
			assert.Equal(t, testCase.Expected.Response.Type, response.Type)
			assert.Equal(t, testCase.Expected.Response.DomainUuid, response.DomainUuid)
			assert.Equal(t, testCase.Expected.Response.Ipa.CaCerts, response.Ipa.CaCerts)
			assert.Equal(t, testCase.Expected.Response.Ipa.Servers, response.Ipa.Servers)
			assert.Equal(t, testCase.Expected.Response.Ipa.RealmName, response.Ipa.RealmName)
		}
	}

}

func TestFillCert(t *testing.T) {
	notValidBefore := time.Now()
	notValidAfter := notValidBefore.Add(time.Hour * 24)
	type TestCaseGiven struct {
		To   *public.DomainIpaCert
		From *model.IpaCert
	}
	type TestCaseExpected struct {
		Err error
		To  public.DomainIpaCert
	}
	type TestCase struct {
		Name     string
		Given    TestCaseGiven
		Expected TestCaseExpected
	}
	testCases := []TestCase{
		{
			Name:  "'to' is nil",
			Given: TestCaseGiven{},
			Expected: TestCaseExpected{
				Err: fmt.Errorf("'to' cannot be nil"),
				To:  public.DomainIpaCert{},
			},
		},
		{
			Name: "'from' is nil",
			Given: TestCaseGiven{
				To: &public.DomainIpaCert{},
			},
			Expected: TestCaseExpected{
				Err: fmt.Errorf("'from' cannot be nil"),
				To:  public.DomainIpaCert{},
			},
		},
		{
			Name: "Success copy",
			Given: TestCaseGiven{
				To: &public.DomainIpaCert{},
				From: &model.IpaCert{
					Nickname:       "MYDOMAIN.EXAMPLE.IPA CA",
					Issuer:         "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE.COM",
					NotValidBefore: notValidBefore,
					NotValidAfter:  notValidAfter,
					SerialNumber:   "1",
					Subject:        "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE.COM",
					Pem:            "-----BEGIN CERTIFICATE-----\nMII...\n-----END CERTIFICATE-----\n",
				},
			},
			Expected: TestCaseExpected{
				Err: nil,
				To: public.DomainIpaCert{
					Nickname:       "MYDOMAIN.EXAMPLE.IPA CA",
					Issuer:         "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE.COM",
					NotValidBefore: notValidBefore,
					NotValidAfter:  notValidAfter,
					SerialNumber:   "1",
					Subject:        "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE.COM",
					Pem:            "-----BEGIN CERTIFICATE-----\nMII...\n-----END CERTIFICATE-----\n",
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Log(testCase.Name)
		// I instantiate directly because the public methods
		// are not part of the interface
		p := domainPresenter{}
		err := p.FillCert(testCase.Given.To, testCase.Given.From)
		if testCase.Expected.Err != nil {
			assert.EqualError(t, err, testCase.Expected.Err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.Expected.To, *testCase.Given.To)
		}
	}
}

func TestFillServer(t *testing.T) {
	type TestCaseGiven struct {
		To   *public.DomainIpaServer
		From *model.IpaServer
	}
	type TestCaseExpected struct {
		Err error
		To  public.DomainIpaServer
	}
	type TestCase struct {
		Name     string
		Given    TestCaseGiven
		Expected TestCaseExpected
	}
	testCases := []TestCase{
		{
			Name:  "'to' is nil",
			Given: TestCaseGiven{},
			Expected: TestCaseExpected{
				Err: fmt.Errorf("'to' cannot be nil"),
				To:  public.DomainIpaServer{},
			},
		},
		{
			Name: "'from' is nil",
			Given: TestCaseGiven{
				To: &public.DomainIpaServer{},
			},
			Expected: TestCaseExpected{
				Err: fmt.Errorf("'from' cannot be nil"),
				To:  public.DomainIpaServer{},
			},
		},
		{
			Name: "Success copy",
			Given: TestCaseGiven{
				To: &public.DomainIpaServer{},
				From: &model.IpaServer{
					FQDN:                "server1.mydomain.example",
					RHSMId:              "547ce70c-9eb5-4783-a619-086aa26f88e5",
					CaServer:            true,
					HCCEnrollmentServer: true,
					HCCUpdateServer:     true,
					PKInitServer:        true,
				},
			},
			Expected: TestCaseExpected{
				Err: nil,
				To: public.DomainIpaServer{
					Fqdn:                  "server1.mydomain.example",
					SubscriptionManagerId: "547ce70c-9eb5-4783-a619-086aa26f88e5",
					CaServer:              true,
					HccEnrollmentServer:   true,
					HccUpdateServer:       true,
					PkinitServer:          true,
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Log(testCase.Name)
		// I instantiate directly because the public methods
		// are not part of the interface
		p := domainPresenter{}
		err := p.FillServer(testCase.Given.To, testCase.Given.From)
		if testCase.Expected.Err != nil {
			assert.EqualError(t, err, testCase.Expected.Err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.Expected.To, *testCase.Given.To)
		}
	}
}

func TestRegisterIpaCaCerts(t *testing.T) {
	var (
		err error
	)
	p := &domainPresenter{}
	notValidBefore := time.Now()
	notValidAfter := notValidBefore.Add(time.Hour * 24)
	ipa := model.Domain{
		IpaDomain: &model.Ipa{
			CaCerts: []model.IpaCert{
				{
					Nickname:       "MYDOMAIN.EXAMPLE.IPA CA",
					Issuer:         "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE.COM",
					NotValidBefore: notValidBefore.UTC(),
					NotValidAfter:  notValidAfter.UTC(),
					SerialNumber:   "1",
					Subject:        "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE.COM",
					Pem:            "-----BEGIN CERTIFICATE-----\nMII...\n-----END CERTIFICATE-----\n",
				},
			},
		},
	}
	ipaNilCerts := model.Domain{
		IpaDomain: &model.Ipa{},
	}
	output := public.DomainResponse{}

	err = p.registerIpaCaCerts(nil, &output)
	assert.NoError(t, err)

	err = p.registerIpaCaCerts(&ipa, nil)
	assert.NoError(t, err)

	err = p.registerIpaCaCerts(&ipaNilCerts, nil)
	assert.NoError(t, err)

	err = p.registerIpaCaCerts(&ipa, &output)
	assert.NoError(t, err)
	assert.Equal(t, ipa.IpaDomain.CaCerts[0].Nickname, output.Ipa.CaCerts[0].Nickname)
	assert.Equal(t, ipa.IpaDomain.CaCerts[0].Issuer, output.Ipa.CaCerts[0].Issuer)
	assert.Equal(t, ipa.IpaDomain.CaCerts[0].NotValidAfter, output.Ipa.CaCerts[0].NotValidAfter)
	assert.Equal(t, ipa.IpaDomain.CaCerts[0].NotValidBefore, output.Ipa.CaCerts[0].NotValidBefore)
	assert.Equal(t, ipa.IpaDomain.CaCerts[0].SerialNumber, output.Ipa.CaCerts[0].SerialNumber)
	assert.Equal(t, ipa.IpaDomain.CaCerts[0].Subject, output.Ipa.CaCerts[0].Subject)
	assert.Equal(t, ipa.IpaDomain.CaCerts[0].Pem, output.Ipa.CaCerts[0].Pem)
}

func TestRegisterIpa(t *testing.T) {
	tokenExpiration := &time.Time{}
	*tokenExpiration = time.Now()
	type TestCaseExpected struct {
		Domain *public.DomainResponse
		Err    error
	}
	type TestCase struct {
		Name     string
		Given    *model.Domain
		Expected TestCaseExpected
	}
	testCases := []TestCase{
		{
			Name:  "ipa is nil",
			Given: nil,
			Expected: TestCaseExpected{
				Domain: &public.DomainResponse{
					Ipa: public.DomainResponseIpa{},
				},
				Err: fmt.Errorf("'ipa' cannot be nil"),
			},
		},
		{
			Name: "Success case",
			Given: &model.Domain{
				Type: pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					RealmName:       pointy.String("mydomain.example"),
					Token:           pointy.String("71ad4978-c768-11ed-ad69-482ae3863d30"),
					TokenExpiration: tokenExpiration,
					RealmDomains:    pq.StringArray{"server.mydomain.example"},
					CaCerts: []model.IpaCert{
						{
							Nickname:     "MYDOMAIN.EXAMPLE IPA CA",
							Issuer:       "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE",
							Subject:      "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE",
							SerialNumber: "1",
							Pem:          "-----BEGIN CERTIFICATE-----\nMII...\n-----END CERTIFICATE-----\n",
						},
					},
					Servers: []model.IpaServer{
						{
							FQDN:                "server.mydomain.example",
							RHSMId:              "c4a80438-c768-11ed-a60e-482ae3863d30",
							PKInitServer:        true,
							CaServer:            true,
							HCCEnrollmentServer: true,
							HCCUpdateServer:     true,
						},
					},
				},
			},
			Expected: TestCaseExpected{
				Domain: &public.DomainResponse{
					Ipa: public.DomainResponseIpa{
						RealmName:       "mydomain.example",
						Token:           nil,
						TokenExpiration: nil,
						RealmDomains:    pq.StringArray{"server.mydomain.example"},
						CaCerts: []public.DomainIpaCert{
							{
								Nickname:     "MYDOMAIN.EXAMPLE IPA CA",
								Issuer:       "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE",
								Subject:      "CN=Certificate Authority,O=MYDOMAIN.EXAMPLE",
								SerialNumber: "1",
								Pem:          "-----BEGIN CERTIFICATE-----\nMII...\n-----END CERTIFICATE-----\n",
							},
						},
						Servers: []public.DomainIpaServer{
							{
								Fqdn:                  "server.mydomain.example",
								SubscriptionManagerId: "c4a80438-c768-11ed-a60e-482ae3863d30",
								PkinitServer:          true,
								CaServer:              true,
								HccEnrollmentServer:   true,
								HccUpdateServer:       true,
							},
						},
					},
				},
				Err: nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Log(testCase.Name)
		p := NewDomainPresenter()
		ipa, err := p.Register(testCase.Given)
		if testCase.Expected.Err != nil {
			assert.EqualError(t, err, testCase.Expected.Err.Error())
			assert.Nil(t, ipa)
		} else {
			assert.NoError(t, err)
			require.NotNil(t, ipa)
			assert.Equal(t, testCase.Expected.Domain.Ipa.RealmName, ipa.Ipa.RealmName)
			require.Equal(t, len(testCase.Expected.Domain.Ipa.RealmDomains), len(ipa.Ipa.RealmDomains))
			for i := range ipa.Ipa.RealmDomains {
				assert.Equal(t, testCase.Expected.Domain.Ipa.RealmDomains[i], ipa.Ipa.RealmDomains[i])
			}
			require.Equal(t, len(testCase.Expected.Domain.Ipa.CaCerts), len(ipa.Ipa.CaCerts))
			for i := range ipa.Ipa.CaCerts {
				assert.Equal(t, testCase.Expected.Domain.Ipa.CaCerts[i], ipa.Ipa.CaCerts[i])
			}
			require.Equal(t, len(testCase.Expected.Domain.Ipa.Servers), len(ipa.Ipa.Servers))
			for i := range ipa.Ipa.Servers {
				assert.Equal(t, testCase.Expected.Domain.Ipa.Servers[i], ipa.Ipa.Servers[i])
			}
		}
	}
}
