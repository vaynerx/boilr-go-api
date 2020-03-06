package app

import (
	bytes "bytes"
	fmt "fmt"

	viper "github.com/spf13/viper"
)

// SkaffoldAPIData defines all available `api` entities.
type SkaffoldAPIData struct {
	main *SkaffoldAPIMain
}

// SkaffoldAPIMain defines configuration settings for the `main` API.
type SkaffoldAPIMain struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemData defines all available `system` entities.
type SkaffoldSystemData struct {
	account *SkaffoldSystemAccount
	auth    *SkaffoldSystemAuth
	media   *SkaffoldSystemMedia
	youtube *SkaffoldSystemYoutube
}

// SkaffoldSystemAccount defines configuration settings for the `account` system.
type SkaffoldSystemAccount struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemAPI defines configuration settings for the `api` system.
type SkaffoldSystemAPI struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemAuth defines configuration settings for the `auth` system.
type SkaffoldSystemAuth struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemMedia defines configuration settings for the `media` system.
type SkaffoldSystemMedia struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemYoutube defines configuration settings for the `youtube` system.
type SkaffoldSystemYoutube struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldClientData defines all available `client` entities.
type SkaffoldClientData struct {
	auth *SkaffoldClientAuth
	main *SkaffoldClientMain
}

// SkaffoldClientAuth defines configuration settings for the `auth` client.
type SkaffoldClientAuth struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldClientMain defines configuration settings for the `main` client.
type SkaffoldClientMain struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldAppData defines all available `app` entities.
type SkaffoldAppData struct {
	hydra      *SkaffoldAppHydra
	postgresql *SkaffoldAppPostgresql
	prisma     *SkaffoldAppPrisma
}

// SkaffoldAppHydra defines configuration settings for the `hydra` application.
type SkaffoldAppHydra struct {
	Host         string
	Port         string
	Address      string
	AdminHost    string
	AdminPort    string
	AdminAddress string
}

// SkaffoldAppPostgresql defines configuration settings for the `postgresql` application.
type SkaffoldAppPostgresql struct {
	Host     string
	Port     string
	Address  string
	Username string
	Password string
}

// SkaffoldAppPrisma defines configuration settings for the `prisma` application.
type SkaffoldAppPrisma struct {
	Host    string
	Port    string
	Address string
	Secret  string
}

var skaffoldConf = `apiVersion: skaffold/v2alpha4
kind: Config
build:
  tagPolicy:
    gitCommit:
      variant: AbbrevTreeSha
profiles:

  - name: development
    build:
      local:
        push: false
      artifacts:
        # --- Clients ---

        - image: client-main
          context: src/clients/main
          sync:
            infer:
              - "frontend/src/**/*"
          docker:
            dockerfile: .development.dockerfile

        - image: client-auth
          context: src/clients/auth
          sync:
            infer:
              - "frontend/src/**/*"
          docker:
            dockerfile: .development.dockerfile

        # --- Services ---

        - image: service-api
          context: src/services/api
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"

        - image: service-auth
          context: src/services/auth
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"              

        - image: service-media
          context: src/services/media
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"
              
        - image: service-account
          context: src/services/account
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"              

        - image: service-youtube
          context: src/services/youtube
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              PYTHON_VERSION: "{{ .PYTHON_VERSION }}"

        # --- Gateways ---

        - image: gateway-auth
          context: src/gateways/auth
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"

        - image: gateway-media
          context: src/gateways/media
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"   

        - image: gateway-account
          context: src/gateways/account
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"               

        - image: gateway-youtube
          context: src/gateways/youtube
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"

    deploy:
      helm:
        releases:
          - name: "config-common"
            chartPath: "k8s/charts/config/common"
            setValueTemplates:
              # --- ENV Clients ---

              env.CLIENT_MAIN_HOST: "{{ .ENV_CLIENT_MAIN_HOST }}"
              env.CLIENT_MAIN_PORT: "{{ .ENV_CLIENT_MAIN_PORT }}"
              env.CLIENT_MAIN_STORYBOOK_PORT: "{{ .ENV_CLIENT_MAIN_STORYBOOK_PORT }}"

              env.CLIENT_AUTH_HOST: "{{ .ENV_CLIENT_AUTH_HOST }}"
              env.CLIENT_AUTH_PORT: "{{ .ENV_CLIENT_AUTH_PORT }}"
              env.CLIENT_AUTH_STORYBOOK_PORT: "{{ .ENV_CLIENT_AUTH_STORYBOOK_PORT }}"

              # --- ENV Services ---

              env.SERVICE_API_SERVICE_HOST: "{{ .ENV_SERVICE_API_SERVICE_HOST }}"
              env.SERVICE_API_SERVICE_PORT: "{{ .ENV_SERVICE_API_SERVICE_PORT }}"

              env.SERVICE_AUTH_SERVICE_HOST: "{{ .ENV_SERVICE_AUTH_SERVICE_HOST }}"
              env.SERVICE_AUTH_SERVICE_PORT: "{{ .ENV_SERVICE_AUTH_SERVICE_PORT }}"
              env.SERVICE_AUTH_SESSION_SECRET: "{{ .ENV_SERVICE_AUTH_SESSION_SECRET }}"

              env.SERVICE_MEDIA_SERVICE_HOST: "{{ .ENV_SERVICE_MEDIA_SERVICE_HOST }}"
              env.SERVICE_MEDIA_SERVICE_PORT: "{{ .ENV_SERVICE_MEDIA_SERVICE_PORT }}"

              env.SERVICE_ACCOUNT_SERVICE_HOST: "{{ .ENV_SERVICE_ACCOUNT_SERVICE_HOST }}"
              env.SERVICE_ACCOUNT_SERVICE_PORT: "{{ .ENV_SERVICE_ACCOUNT_SERVICE_PORT }}"

              env.SERVICE_YOUTUBE_SERVICE_PORT: "{{ .ENV_SERVICE_YOUTUBE_SERVICE_PORT }}"
              env.SERVICE_YOUTUBE_SERVICE_HOST: "{{ .ENV_SERVICE_YOUTUBE_SERVICE_HOST }}"

              # --- ENV Gateways ---

              env.GATEWAY_AUTH_SERVICE_HOST: "{{ .ENV_GATEWAY_AUTH_SERVICE_HOST }}"
              env.GATEWAY_AUTH_PORT: "{{ .ENV_GATEWAY_AUTH_PORT }}"

              env.GATEWAY_MEDIA_HOST: "{{ .ENV_GATEWAY_MEDIA_HOST }}"
              env.GATEWAY_MEDIA_PORT: "{{ .ENV_GATEWAY_MEDIA_PORT }}"

              env.GATEWAY_ACCOUNT_SERVICE_HOST: "{{ .ENV_GATEWAY_ACCOUNT_SERVICE_HOST }}"
              env.GATEWAY_ACCOUNT_SERVICE_PORT: "{{ .ENV_GATEWAY_ACCOUNT_SERVICE_PORT }}"

              env.GATEWAY_YOUTUBE_HOST: "{{ .ENV_GATEWAY_YOUTUBE_HOST }}"
              env.GATEWAY_YOUTUBE_PORT: "{{ .ENV_GATEWAY_YOUTUBE_PORT }}"                            

              # --- ENV Apps ---

              env.APP_NGINX_INGRESS_HOST: "{{ .ENV_APP_NGINX_INGRESS_HOST }}"
              env.APP_NGINX_INGRESS_PORT_HTTP: "{{ .ENV_APP_NGINX_INGRESS_PORT_HTTP }}"
              env.APP_NGINX_INGRESS_PORT_HTTPS: "{{ .ENV_APP_NGINX_INGRESS_PORT_HTTPS }}"
              env.APP_NGINX_INGRESS_LOAD_BALANCER_IP: "{{ .ENV_APP_NGINX_INGRESS_LOAD_BALANCER_IP }}"

              env.APP_POSTGRESQL_PORT: "{{ .ENV_APP_POSTGRESQL_PORT }}"
              env.APP_POSTGRESQL_HOST: "{{ .ENV_APP_POSTGRESQL_HOST }}"
              env.APP_POSTGRESQL_USER: "{{ .ENV_APP_POSTGRESQL_USER }}"
              env.APP_POSTGRESQL_PASSWORD: "{{ .ENV_APP_POSTGRESQL_PASSWORD }}"

              env.APP_HYDRA_HOST: "{{ .ENV_APP_HYDRA_HOST }}"
              env.APP_HYDRA_PORT: "{{ .ENV_APP_HYDRA_PORT }}"
              env.APP_HYDRA_ADMIN_PORT: "{{ .ENV_APP_HYDRA_ADMIN_PORT }}"

              env.APP_PRISMA_HOST: "{{ .ENV_APP_PRISMA_HOST }}"
              env.APP_PRISMA_PORT: "{{ .ENV_APP_PRISMA_PORT }}"
              env.APP_PRISMA_SECRET: "{{ .ENV_APP_PRISMA_SECRET }}"

              # --- ENV Misc. ---

              env.PROJECT_STAGE: "{{ .ENV_PROJECT_STAGE }}"

          # --- Apps ---

          - name: "app-nginx-ingress"
            chartPath: stable/nginx-ingress
            version: "1.29.1"
            remote: true
            setValueTemplates:
              defaultBackend.enabled: true
              controller.publishService.enabled: true
              controller.service.ports.http: "{{ .ENV_APP_NGINX_INGRESS_PORT_HTTP }}"
              controller.service.ports.https: "{{ .ENV_APP_NGINX_INGRESS_PORT_HTTPS }}"

          - name: "app-cert-manager"
            chartPath: jetstack/cert-manager
            version: "v0.12.0"
            remote: true
            wait: true
            namespace: "cert-manager"
            setValueTemplates:
              webhook.enabled: false
              ingressShim.defaultIssuerName: "{{ .ENV_PROJECT_STAGE }}"
              ingressShim.defaultIssuerKind: ClusterIssuer

          - name: "app-prisma"
            chartPath: stable/prisma
            version: "1.2.1"
            remote: true
            setValueTemplates:
              service.port: "{{ .ENV_APP_PRISMA_PORT }}"
              auth.enabled: true
              auth.secret: "{{ .ENV_APP_PRISMA_SECRET }}"
              database.user: "{{ .ENV_APP_POSTGRESQL_USER }}"  
              database.password: "{{ .ENV_APP_POSTGRESQL_PASSWORD }}"  
              postgresql.enabled: true
              postgresql.service.port: "{{ .ENV_APP_POSTGRESQL_PORT }}"
              postgresql.postgresPassword: "{{ .ENV_APP_POSTGRESQL_PASSWORD }}"
              postgresql.postgresUser: "{{ .ENV_APP_POSTGRESQL_USER }}"                    

          - name: "app-hydra"
            chartPath: ory/hydra
            version: "0.0.48"
            remote: true
            setValueTemplates:
              hydra.config.secrets.system: "{{ .ENV_HYDRA_SECRET }}"
              hydra.config.dsn: "postgres://{{ .ENV_APP_POSTGRESQL_USER }}:{{ .ENV_APP_POSTGRESQL_PASSWORD }}@{{ .ENV_APP_POSTGRESQL_HOST }}:{{ .ENV_APP_POSTGRESQL_PORT }}/prisma?sslmode=disable"
              hydra.config.urls.self.issuer: "http://hydra.se.development.teamgaryvee.com/"
              hydra.config.urls.login: "http://auth.se.development.teamgaryvee.com/login"
              hydra.config.urls.consent: "http://auth.se.development.teamgaryvee.com/consent"
              hydra.config.urls.logout: "http://auth.se.development.teamgaryvee.com/logout"
              ingress.public.enabled: true
              ingress.admin.enabled: true
              hydra.dangerousForceHttp: true
              hydra.autoMigrate: true
            overrides:
              hydra.dangerousAllowInsecureRedirectUrls:
                - "http://hydra.se.development.teamgaryvee.com"
                - "https://hydra.se.development.teamgaryvee.com"
                - "127.0.0.1/32"

          - name: "app-kafka"
            chartPath: incubator/kafka      
            version: "0.20.8"
            remote: true
              

          # --- Services ---

          - name: "service-api"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-api
            setValueTemplates:
              service.port: "{{ .ENV_SERVICE_API_SERVICE_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}
            overrides:
              ingress:
                enabled: true
                annotations:
                  kubernetes.io/ingress.class: nginx
                  nginx.ingress.kubernetes.io/use-regex: "true"
                hosts:
                  - host: api.se.development.teamgaryvee.com
                    paths:
                      - /.*

          - name: "service-auth"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-auth
            setValueTemplates:
              service.port: "{{ .ENV_SERVICE_AUTH_SERVICE_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}       
              
          - name: "service-media"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-media
            setValueTemplates:
              service.port: "{{ .ENV_SERVICE_MEDIA_SERVICE_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}                 

          - name: "service-account"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-account
            setValueTemplates:
              service.port: "{{ .ENV_SERVICE_ACCOUNT_SERVICE_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}                 
                      
          - name: "service-youtube"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-youtube
            setValueTemplates:
              service.port: "{{ .ENV_SERVICE_YOUTUBE_SERVICE_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}                      

          # --- Gateways ---

          - name: "gateway-auth"
            chartPath: "k8s/charts/lib/gateway"
            values:
              image: gateway-auth
            setValueTemplates:
              service.port: "{{ .ENV_GATEWAY_AUTH_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}

          - name: "gateway-youtube"
            chartPath: "k8s/charts/lib/gateway"
            values:
              image: gateway-youtube
            setValueTemplates:
              service.port: "{{ .ENV_GATEWAY_YOUTUBE_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}     

          # --- Clients ---

          - name: "client-main"
            chartPath: "k8s/charts/lib/client"
            values:
              image: client-main
            setValueTemplates:
              service.port: "{{ .ENV_CLIENT_MAIN_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}
            overrides:
              ingress:
                enabled: true
                annotations:
                  kubernetes.io/ingress.class: nginx
                  nginx.ingress.kubernetes.io/use-regex: "true"
                hosts:
                  - host: se.development.teamgaryvee.com
                    paths:
                      - /.*

          - name: "client-main-storybook"
            chartPath: "k8s/charts/lib/client"
            values:
              image: client-main
            setValueTemplates:
              service.port: "{{ .ENV_CLIENT_MAIN_STORYBOOK_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}
            overrides:
              command: ["npm"]
              args: ["run", "storybook"]
              ingress:
                enabled: true
                annotations:
                  kubernetes.io/ingress.class: nginx
                  nginx.ingress.kubernetes.io/use-regex: "true"
                hosts:
                  - host: storybook.se.development.teamgaryvee.com
                    paths:
                      - /.*

          - name: "client-auth"
            chartPath: "k8s/charts/lib/client"
            values:
              image: client-auth
            setValueTemplates:
              service.port: "{{ .ENV_CLIENT_MAIN_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}
            overrides:
              ingress:
                enabled: true
                annotations:
                  kubernetes.io/ingress.class: nginx
                  nginx.ingress.kubernetes.io/use-regex: "true"
                hosts:
                  - host: auth.se.development.teamgaryvee.com
                    paths:
                      - /.*

          - name: "client-auth-storybook"
            chartPath: "k8s/charts/lib/client"
            values:
              image: client-auth
            setValueTemplates:
              service.port: "{{ .ENV_CLIENT_AUTH_STORYBOOK_PORT }}"
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}
            overrides:
              command: ["npm"]
              args: ["run", "storybook"]
              ingress:
                enabled: true
                annotations:
                  kubernetes.io/ingress.class: nginx
                  nginx.ingress.kubernetes.io/use-regex: "true"
                hosts:
                  - host: storybook.auth.se.development.teamgaryvee.com
                    paths:
                      - /.*

          - name: "client-linkerd-dashboard"
            chartPath: "k8s/charts/lib/linkerd/linkerd-dashboard"
            namespace: linkerd
            setValueTemplates:
              service.port: "{{ .ENV_CLIENT_LINKERD_DASHBOARD_PORT }}"
              service.name: linkerd-web
            valuesFiles:
              - ./config/settings/globals.yaml
            overrides:
              ingress:
                enabled: true
                annotations:
                  kubernetes.io/ingress.class: nginx
                  nginx.ingress.kubernetes.io/use-regex: "true"
                  nginx.ingress.kubernetes.io/configuration-snippet: |
                    proxy_set_header l5d-dst-override $service_name.$namespace.svc.cluster.local:8084;
                    proxy_set_header Origin "";
                    proxy_hide_header l5d-remote-ip;
                    proxy_hide_header l5d-server-id;
                  nginx.ingress.kubernetes.io/upstream-vhost: $service_name.$namespace.svc.cluster.local:8084
                  nginx.ingress.kubernetes.io/auth-type: basic
                  nginx.ingress.kubernetes.io/auth-secret: web-ingress-auth
                  nginx.ingress.kubernetes.io/auth-realm: "Authentication Required"
                hosts:
                  - host: linkerd.se.development.teamgaryvee.com
                    paths:
                      - /.*

          - name: "config-cert-manager"
            chartPath: "k8s/charts/config/cert-manager"
            wait: true
            setValueTemplates:
              letsencrypt.issuer: "{{ .ENV_PROJECT_STAGE }}"
              letsencrypt.email: "{{ .ENV_LETSENCRYPT_EMAIL }}"
              letsencrypt.production.enabled: false

  - name: api
    build:
      local:
        push: false
      artifacts:

        # --- Services ---

        - image: service-api
          context: src/services/api
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"

        - image: service-auth
          context: src/services/auth
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"              

        - image: service-media
          context: src/services/media
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"
              
        - image: service-account
          context: src/services/account
          sync:
            infer:
              - "api/**/*"
              - "cmd/**/*"
              - "pkg/**/*"
              - "util/**/*"
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              GOLANG_VERSION: "{{ .GOLANG_VERSION }}"              

        - image: service-youtube
          context: src/services/youtube
          docker:
            dockerfile: .development.dockerfile
            buildArgs:
              PYTHON_VERSION: "{{ .PYTHON_VERSION }}"

    deploy:
      helm:
        releases:
          - name: "config-common"
            chartPath: "k8s/charts/config/common"
            setValueTemplates:          
              env.POSTGRESQL_ADDRESS: "{{ .POSTGRESQL_ADDRESS }}"
              env.POSTGRESQL_DATABASE: "{{ .POSTGRESQL_DATABASE }}"
              env.POSTGRESQL_USER: "{{ .POSTGRESQL_USER }}"
              env.POSTGRESQL_PASSWORD: "{{ .POSTGRESQL_PASSWORD }}"
              env.PRISMA_SECRET: "{{ .PRISMA_SECRET }}"
              env.HYDRA_SECRET: "{{ .HYDRA_SECRET }}"

          # --- Apps ---

          - name: "app-nginx-ingress"
            chartPath: stable/nginx-ingress
            version: "1.29.1"
            remote: true
            setValueTemplates:
              defaultBackend.enabled: "false"
              controller.publishService.enabled: "true"
              controller.service.ports.http: 80
              controller.service.ports.https: 443

          - name: "app-prisma"
            chartPath: stable/prisma
            version: "1.2.1"
            remote: true
            setValueTemplates:
              auth.enabled: "true"
              auth.secret: "{{ .PRISMA_SECRET }}"
              database.user: "{{ .POSTGRESQL_USER }}"  
              database.password: "{{ .POSTGRESQL_PASSWORD }}"  
              postgresql.enabled: "true"
              postgresql.postgresPassword: "{{ .POSTGRESQL_PASSWORD }}"
              postgresql.postgresUser: "{{ .POSTGRESQL_USER }}"                    

          - name: "app-hydra"
            chartPath: ory/hydra
            version: "0.0.48"
            remote: true
            setValueTemplates:
              ingress.public.enabled: "true"
              ingress.admin.enabled: "true"
              hydra.dangerousForceHttp: "true"
              hydra.autoMigrate: "true"
              hydra.config.secrets.system: "{{ .HYDRA_SECRET }}"
              hydra.config.dsn: "postgres://{{ .POSTGRESQL_USER }}:{{ .POSTGRESQL_PASSWORD }}@{{ .POSTGRESQL_ADDRESS }}/{{ .POSTGRESQL_DATABASE }}?sslmode=disable"
              hydra.config.urls.self.issuer: "http://hydra.se.development.teamgaryvee.com/"
              hydra.config.urls.login: "http://auth.se.development.teamgaryvee.com/login"
              hydra.config.urls.consent: "http://auth.se.development.teamgaryvee.com/consent"
              hydra.config.urls.logout: "http://auth.se.development.teamgaryvee.com/logout"

          - name: "app-kafka"
            chartPath: incubator/kafka      
            version: "0.20.8"
            remote: true

          # --- Services ---

          - name: "service-api"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-api
            setValueTemplates:
              service.port: 7000
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}
            overrides:
              ingress:
                enabled: true
                annotations:
                  kubernetes.io/ingress.class: nginx
                  nginx.ingress.kubernetes.io/use-regex: "true"
                hosts:
                  - host: api.se.development.teamgaryvee.com
                    paths:
                      - /.*

          - name: "service-auth"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-auth
            setValueTemplates:
              service.port: 7010
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}       
              
          - name: "service-media"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-media
            setValueTemplates:
              service.port: 7020
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}                 

          - name: "service-account"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-account
            setValueTemplates:
              service.port: 7030
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}                 
                      
          - name: "service-youtube"
            chartPath: "k8s/charts/lib/service"
            values:
              image: service-youtube
            setValueTemplates:
              service.port: 7040
            valuesFiles:
              - ./config/settings/globals.yaml
            imageStrategy:
              helm: {}`

type profiles map[string]interface{}
type releases map[string]interface{}

func getSkaffold() {
	confBuf := &bytes.Buffer{}
	confBuf.WriteString(skaffoldConf)

	profiles := make([]profiles, 0)

	mainViper := viper.New()
	mainViper.SetConfigType("yaml")
	_ = mainViper.ReadConfig(confBuf)
	_ = mainViper.UnmarshalKey("profiles", &profiles)

	fmt.Println(profiles[0]["deploy"]["helm"]["releases"])
}
