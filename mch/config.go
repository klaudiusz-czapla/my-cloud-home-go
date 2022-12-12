package mch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MchConfig struct {
	Data struct {
		ConfigurationID     string                            `json:"configurationId"`
		ComponentMapUntyped map[string]map[string]interface{} `json:"componentMap"`
		ComponentMap        struct {
			ComWdcAppmanager struct {
				AppstoreURL string `json:"appstore.url"`
			} `json:"com.wdc.appmanager"`
			ComWdDevice struct {
				P2PEnabled    bool `json:"p2p_enabled"`
				RestoreConfig bool `json:"restore_config"`
			} `json:"com.wd.device"`
			ComWdNotification struct {
				TestDouble          float64 `json:"test_double"`
				ReceiveNotification string  `json:"receive_notification"`
				PollingIntervalSec  int     `json:"polling_interval_sec"`
			} `json:"com.wd.notification"`
			ComWdAnalytics struct {
				DeviceLogIntervalSec int    `json:"device_log_interval_sec"`
				RequiredDataScrub    bool   `json:"required_data_scrub"`
				StartLogging         string `json:"start_logging"`
				WddirectLog          int    `json:"wddirect_log"`
			} `json:"com.wd.analytics"`
			CloudServiceUrlsUntyped map[string]interface{} `json:"cloud.service.urls"`
			CloudServiceUrls        struct {
				ServiceAccountLoginURL                string `json:"service.account.login.url"`
				ServiceAccountRedirectURL             string `json:"service.account.redirect.url"`
				ServiceDiscoveryHelpURL               string `json:"service.discovery.help.url"`
				ServiceIbidiscoveryHelpURL            string `json:"service.ibidiscovery.help.url"`
				ServiceDiscoveryMycloudURL            string `json:"service.discovery.mycloud.url"`
				ServiceDiscoveryOlcURL                string `json:"service.discovery.olc.url"`
				ServiceDiscoveryPrivacypolicyURL      string `json:"service.discovery.privacypolicy.url"`
				ServiceDiscoveryUsermanualURL         string `json:"service.discovery.usermanual.url"`
				ServiceLocalyticsproxyProfileapiURL   string `json:"service.localyticsproxy.profileapi.url"`
				ServiceDeviceURL                      string `json:"service.device.url"`
				ServiceDevicenetworkURL               string `json:"service.devicenetwork.url"`
				ServiceDevicenotificationURL          string `json:"service.devicenotification.url"`
				AnalyticsClientDesktopMacOSURL        string `json:"analytics.client.desktop.macOS.url"`
				AnalyticsClientOdaURL                 string `json:"analytics.client.oda.url"`
				AnalyticsClientWebURL                 string `json:"analytics.client.web.url"`
				AnalyticsClientErwebappURL            string `json:"analytics.client.erwebapp.url"`
				AnalyticsClientKeplerwebappURL        string `json:"analytics.client.keplerwebapp.url"`
				ServiceDasURL                         string `json:"service.das.url"`
				ServiceProductURL                     string `json:"service.product.url"`
				ServiceRentalURL                      string `json:"service.rental.url"`
				ServiceKasURL                         string `json:"service.kas.url"`
				ServiceIotServicesURL                 string `json:"service.iotServices.url"`
				ServiceCognitoPoolURL                 string `json:"service.cognito.pool.url"`
				ServiceCognitoPoolDomain              string `json:"service.cognito.pool.domain"`
				ServiceTsmURL                         string `json:"service.tsm.url"`
				ServiceProxyURL                       string `json:"service.proxy.url"`
				ServiceOtaURL                         string `json:"service.ota.url"`
				AnalyticsURL                          string `json:"analytics.url"`
				WebclientNewSessionURL                string `json:"webclient.new_session.url"`
				WebclientNewSessionURLIbi             string `json:"webclient.new_session.url.ibi"`
				WebclientNewSessionURLMch             string `json:"webclient.new_session.url.mch"`
				ServiceFeebackserviceURL              string `json:"service.feebackservice.url"`
				ServiceAuth0URL                       string `json:"service.auth0.url"`
				ServiceUpthereURL                     string `json:"service.upthere.url"`
				ServiceM2MURL                         string `json:"service.m2m.url"`
				ServiceSwupdateURL                    string `json:"service.swupdate.url"`
				ServiceEventsURL                      string `json:"service.events.url"`
				ServiceDevportalURL                   string `json:"service.devportal.url"`
				ServiceInvitationsURL                 string `json:"service.invitations.url"`
				ServiceThirdpartyfederationURL        string `json:"service.thirdpartyfederation.url"`
				AnalyticsClientDesktopWindowsURL      string `json:"analytics.client.desktop.windows.url"`
				ServiceCommunicationWsURL             string `json:"service.communication.ws.url"`
				AnalyticsClientDesktopCbfsLicenseURL  string `json:"analytics.client.desktop.cbfsLicense.url"`
				AnalyticsDeviceURL                    string `json:"analytics.device.url"`
				AnalyticsDeviceAdminuiURL             string `json:"analytics.device.adminui.url"`
				ServiceAuthURL                        string `json:"service.auth.url"`
				AnalyticsDeviceYodaplusURL            string `json:"analytics.device.yodaplus.url"`
				AnalyticsDeviceYodaplus2URL           string `json:"analytics.device.yodaplus2.url"`
				ServiceLogdumpURL                     string `json:"service.logdump.url"`
				AnalyticsDeviceYodaURL                string `json:"analytics.device.yoda.url"`
				WebclientMycloudURL                   string `json:"webclient.mycloud.url"`
				WebclientMycloudURLIbi                string `json:"webclient.mycloud.url.ibi"`
				WebclientMycloudURLMch                string `json:"webclient.mycloud.url.mch"`
				ServiceConfigURL                      string `json:"service.config.url"`
				ServiceCommunicationURL               string `json:"service.communication.url"`
				AnalyticsDeviceMonarchURL             string `json:"analytics.device.monarch.url"`
				AnalyticsDeviceMonarch2URL            string `json:"analytics.device.monarch2.url"`
				ServiceShareURL                       string `json:"service.share.url"`
				ServiceSharedfilesURL                 string `json:"service.sharedfiles.url"`
				ServiceNotificationURL                string `json:"service.notification.url"`
				ServiceNotifyserviceURL               string `json:"service.notifyservice.url"`
				ServicePortalURL                      string `json:"service.portal.url"`
				WebclientSignupURL                    string `json:"webclient.signup.url"`
				WebclientSignupURLIbi                 string `json:"webclient.signup.url.ibi"`
				WebclientSignupURLMch                 string `json:"webclient.signup.url.mch"`
				AnalyticsDevicePelicanURL             string `json:"analytics.device.pelican.url"`
				AnalyticsDevicePelican2URL            string `json:"analytics.device.pelican2.url"`
				AnalyticsClientMobileURL              string `json:"analytics.client.mobile.url"`
				ServiceTinyconfigURL                  string `json:"service.tinyconfig.url"`
				ServiceAppcatalogURL                  string `json:"service.appcatalog.url"`
				ServiceSumoproxyURL                   string `json:"service.sumoproxy.url"`
				ServiceLocalyticsproxyEventapiURL     string `json:"service.localyticsproxy.eventapi.url"`
				ServicePluginsURL                     string `json:"service.plugins.url"`
				ServicePluginsFileJSON                string `json:"service.plugins.file.json"`
				WebclientForgotPasswordURL            string `json:"webclient.forgot_password.url"`
				WebclientIbiForgotPasswordURL         string `json:"webclient.ibi.forgot_password.url"`
				ServiceClientURL                      string `json:"service.client.url"`
				ServiceIbiHome                        string `json:"service.ibi.home"`
				ServiceIotURL                         string `json:"service.iot.url"`
				ServiceMqttURL                        string `json:"service.mqtt.url"`
				WebclientIbiBrowseURL                 string `json:"webclient.ibi.browse.url"`
				AnalyticsDeviceWDMyCloudURL           string `json:"analytics.device.WDMyCloud.url"`
				AnalyticsDeviceWDMyCloudMirrorURL     string `json:"analytics.device.WDMyCloudMirror.url"`
				AnalyticsDeviceMyCloudEX2UltraURL     string `json:"analytics.device.MyCloudEX2Ultra.url"`
				AnalyticsDeviceWDMyCloudEX4100URL     string `json:"analytics.device.WDMyCloudEX4100.url"`
				AnalyticsDeviceMyCloudPR2100URL       string `json:"analytics.device.MyCloudPR2100.url"`
				AnalyticsDeviceMyCloudPR4100URL       string `json:"analytics.device.MyCloudPR4100.url"`
				AnalyticsDeviceWDCloudURL             string `json:"analytics.device.WDCloud.url"`
				AnalyticsDeviceSequoiaURL             string `json:"analytics.device.sequoia.url"`
				AnalyticsDeviceWDMyCloudDL2100URL     string `json:"analytics.device.WDMyCloudDL2100.url"`
				AnalyticsDeviceWDMyCloudDL4100URL     string `json:"analytics.device.WDMyCloudDL4100.url"`
				AnalyticsDeviceWDMyCloudEX2100URL     string `json:"analytics.device.WDMyCloudEX2100.url"`
				AnalyticsDeviceWDMyCloudEX2URL        string `json:"analytics.device.WDMyCloudEX2.url"`
				AnalyticsDeviceWDMyCloudEX4URL        string `json:"analytics.device.WDMyCloudEX4.url"`
				AnalyticsDeviceWDMyCloudMirrorGen1URL string `json:"analytics.device.WDMyCloudMirrorGen1.url"`
				AnalyticsDeviceDarkwingURL            string `json:"analytics.device.darkwing.url"`
				AnalyticsDeviceRocketURL              string `json:"analytics.device.rocket.url"`
				AnalyticsDeviceDraxURL                string `json:"analytics.device.drax.url"`
				AnalyticsEventsURL                    string `json:"analytics.events.url"`
				ServiceEpochURL                       string `json:"service.epoch.url"`
				ServiceSubscriptionURL                string `json:"service.subscription.url"`
			} `json:"cloud.service.urls"`
			ComWdPortal struct {
				PortalAuth0Client          string `json:"portal.auth0.client"`
				PortalAuth0ClientUniversal string `json:"portal.auth0.client.universal"`
			} `json:"com.wd.portal"`
			ComWdConnection struct {
				EnableNotification bool   `json:"enable_notification"`
				AllowRedirect      bool   `json:"allow_redirect"`
				MinCacheSizeMb     string `json:"min_cache_size_mb"`
				CachingThreshold   string `json:"caching_threshold"`
				MaxCacheSizeMb     string `json:"max_cache_size_mb"`
			} `json:"com.wd.connection"`
			ComSandiskIxpandcharger struct {
				FodAttributes struct {
					Android struct {
						FodConfigVersion string        `json:"fod_config_version"`
						AndroidAll       string        `json:"android_all"`
						MakeList         []interface{} `json:"make_list"`
						ModelList        []struct {
							Make     string   `json:"make"`
							Model    string   `json:"model"`
							DwModels []string `json:"dw_models"`
						} `json:"model_list"`
					} `json:"android"`
					Ios struct {
						FodConfigVersion string `json:"fod_config_version"`
						IosAll           string `json:"ios_all"`
						ModelList        []struct {
							PublicModel   string   `json:"public_model"`
							PlatformModel string   `json:"platform_model"`
							DwModels      []string `json:"dw_models"`
						} `json:"model_list"`
					} `json:"ios"`
				} `json:"fod_attributes"`
			} `json:"com.sandisk.ixpandcharger"`
		} `json:"componentMap"`
	} `json:"data"`
}

const (
	configURL = "https://config.mycloud.com/config/v1/config"
)

func GetConfiguration() (*MchConfig, error) {
	resp, err := http.Get(configURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytesArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var config MchConfig
	err = json.Unmarshal(respBytesArr, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *MchConfig) GetCloudServiceUrl(name string) (string, error) {
	if val, ok := c.Data.ComponentMap.CloudServiceUrlsUntyped[name]; ok {
		if s, ok := val.(string); ok {
			return s, nil
		}

		return "", fmt.Errorf("received cloud service url with name '%s' is not a string", name)
	}

	return "", fmt.Errorf("cloud service url with name '%s' were not found", name)
}
