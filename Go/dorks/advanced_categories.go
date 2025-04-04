package dorks

type DorkCategories struct {
    Categories map[string][]string
}

func NewDorkCategories() *DorkCategories {
    return &DorkCategories{
        Categories: map[string][]string{
            "sensitive_files": {
                `site:{domain} ext:env OR ext:yml OR ext:yaml OR ext:conf`,
                `site:{domain} filename:config OR filename:settings`,
                `site:{domain} filename:id_rsa OR filename:id_dsa`,
                `site:{domain} filename:backup OR filename:dump`,
            },
            "exposed_panels": {
                `site:{domain} inurl:admin OR inurl:console`,
                `site:{domain} inurl:jenkins OR inurl:gitlab`,
                `site:{domain} inurl:grafana OR inurl:kibana`,
                `site:{domain} inurl:phpmyadmin OR inurl:wp-admin`,
            },
            "api_endpoints": {
                `site:{domain} inurl:api swagger OR "api documentation"`,
                `site:{domain} inurl:graphql OR inurl:graphiql`,
                `site:{domain} intext:"api_key" OR intext:"apikey"`,
                `site:{domain} inurl:v1 OR inurl:v2 filetype:json`,
            },
            "cloud_resources": {
                `site:{domain} "amazonaws.com" OR "blob.core.windows.net"`,
                `site:{domain} "storage.googleapis.com" OR "firebaseio.com"`,
                `site:{domain} "s3.amazonaws.com" bucket OR upload`,
                `site:{domain} "azurewebsites.net" OR "cloudapp.net"`,
            },
            "security_misconfigs": {
                `site:{domain} "Index of /" password OR admin OR user`,
                `site:{domain} "MYSQL_ROOT" OR "POSTGRESQL_ROOT"`,
                `site:{domain} "BEGIN RSA PRIVATE KEY" OR "BEGIN SSH KEY"`,
                `site:{domain} "phpinfo()" OR "debug=true"`,
            },
        },
    }
}

