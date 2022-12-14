module github.com/wangxn2015/ran-simulator

go 1.16

require (
	github.com/Shopify/sarama v1.37.2 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.2.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/onosproject/onos-api/go v0.9.43
	//github.com/wangxn2015/helmit v1.6.20
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go v0.8.7
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go v0.8.7
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc v0.8.24
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go v0.8.7
	github.com/onosproject/onos-e2t v0.11.7
	github.com/onosproject/onos-lib-go v0.8.17
	github.com/onosproject/onos-test v0.6.4
	github.com/pmcxs/hexgrid v0.0.0-20190126214921-42796ac894ab
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.8.0
	github.com/wangxn2015/helmit v1.6.21-0.20221211101934-72a55699c433
	github.com/wangxn2015/onos-e2t v0.11.1-0.20221213104809-9e750336d72d
	github.com/wangxn2015/onos-lib-go v0.8.16-0.20221213045740-e38a2ad92701
	github.com/wangxn2015/onos-ric-sdk-go v0.8.9-0.20221212153731-3320d18f773e
	github.com/wangxn2015/rrm-son-lib v1.0.4-0.20221211101328-ae037e9ab599
	//github.com/wangxn2015/onos-e2t v0.0.0-00010101000000-000000000000
	//github.com/wangxn2015/onos-lib-go v1.8.15
	//github.com/wangxn2015/onos-ric-sdk-go v0.0.0-00010101000000-000000000000
	//github.com/wangxn2015/rrm-son-lib v1.0.2
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.1
	googlemaps.github.io/maps v1.3.2
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible
)

replace helm.sh/helm/v3 => helm.sh/helm/v3 v3.9.3

//replace github.com/wangxn2015/onos-e2t => /home/baicells/go_project/modified-onos-module/onos-e2t

//replace github.com/wangxn2015/onos-lib-go => /home/baicells/go_project/modified-onos-module/onos-lib-go

//replace github.com/wangxn2015/rrm-son-lib => /home/baicells/go_project/modified-onos-module/rrm-son-lib

//replace github.com/wangxn2015/helmit => /home/baicells/go_project/modified-onos-module/helmit

//replace github.com/wangxn2015/onos-ric-sdk-go => /home/baicells/go_project/modified-onos-module/onos-ric-sdk-go

//-------------
//replace github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_kpm_v2_go => /home/baicells/go_project/modified-onos-module/onos-e2-sm/servicemodels/e2sm_kpm_v2_go
//
//replace github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_mho_go => /home/baicells/go_project/modified-onos-module/onos-e2-sm/servicemodels/e2sm_mho_go
//
//replace github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc => /home/baicells/go_project/modified-onos-module/onos-e2-sm/servicemodels/e2sm_rc
//
//replace github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc_pre_go => /home/baicells/go_project/modified-onos-module/onos-e2-sm/servicemodels/e2sm_rc_pre_go

//-------------

replace github.com/onosproject/onos-api/go v0.9.29 => github.com/onosproject/onos-api/go v0.9.18

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200229013735-71373c6105e3

replace github.com/pmcxs/hexgrid v0.0.0-20190126214921-42796ac894ab => github.com/SeanCondon/hexgrid v0.0.0-20200424141352-c3819a378a18
