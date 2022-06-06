package controllers

import (
	"encoding/json"
	"gitlab-webhook/config"
	"gitlab-webhook/dto"
	"testing"
)

func Test_preDeployment(t *testing.T) {
	type args struct {
		gitlabConfig *config.GitlabConfig
		webHookData  *dto.GitlabWebhooks
	}
	conf, _ := config.Get("oasis-ottopay")
	webhookData := new(dto.GitlabWebhooks)
	json.Unmarshal([]byte(`{"object_kind":"push","event_name":"push","before":"01f50693f832e68546cd1f37c89f8b987483311f","after":"0bc9e055113cf0c7b09b5140ac7a05d9269fdd61","ref":"refs/heads/development","checkout_sha":"0bc9e055113cf0c7b09b5140ac7a05d9269fdd61","message":null,"user_id":119,"user_name":"zaky yudha","user_username":"zaky","user_email":"","user_avatar":"https://andromeda.ottopay.id/uploads/-/system/user/avatar/119/avatar.png","project_id":221,"project":{"id":221,"name":"ottopay-auth-services","description":"","web_url":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services","avatar_url":null,"git_ssh_url":"git@andromeda.ottopay.id:Ottopay/ottopay-auth-services.git","git_http_url":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services.git","namespace":"Ottopay","visibility_level":0,"path_with_namespace":"Ottopay/ottopay-auth-services","default_branch":"master","ci_config_path":null,"homepage":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services","url":"git@andromeda.ottopay.id:Ottopay/ottopay-auth-services.git","ssh_url":"git@andromeda.ottopay.id:Ottopay/ottopay-auth-services.git","http_url":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services.git"},"commits":[{"id":"35aa5694d1bac46afa1dccb88f9a3eb4bd1369bc","message":"OTTPAY-1928 | Handling opsi otp get from rose\n","title":"OTTPAY-1928 | Handling opsi otp get from rose","timestamp":"2021-09-02T20:59:41+07:00","url":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services/-/commit/35aa5694d1bac46afa1dccb88f9a3eb4bd1369bc","author":{"name":"Zaky","email":"zaky.yudha@ottodigital.id"},"added":["helpers/merchant.go"],"modified":["db/users.go","db/users_test.go","models/rosemerchantmodels/merchant_profile.go"],"removed":[]},{"id":"0bc9e055113cf0c7b09b5140ac7a05d9269fdd61","message":"resolve conflict with development\n","title":"resolve conflict with development","timestamp":"2021-09-02T21:11:51+07:00","url":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services/-/commit/0bc9e055113cf0c7b09b5140ac7a05d9269fdd61","author":{"name":"Zaky","email":"zaky.yudha@ottodigital.id"},"added":["helpers/merchant.go"],"modified":["db/users.go","db/users_test.go","models/rosemerchantmodels/merchant_profile.go"],"removed":[]}],"total_commits_count":2,"push_options":{},"repository":{"name":"ottopay-auth-services","url":"git@andromeda.ottopay.id:Ottopay/ottopay-auth-services.git","description":"","homepage":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services","git_http_url":"https://andromeda.ottopay.id/Ottopay/ottopay-auth-services.git","git_ssh_url":"git@andromeda.ottopay.id:Ottopay/ottopay-auth-services.git","visibility_level":0}}`), &webhookData)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test pre-webhook",
			args:args{
				gitlabConfig: conf,
				webHookData:  webhookData,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preDeployment(tt.args.gitlabConfig, tt.args.webHookData)
		})
	}
}