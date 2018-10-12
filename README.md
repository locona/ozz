# ozz

## oathkeeper

### Authenticators
* credential の認証を担当.
* http requestの検証やsubjectIDの認証を行う.
* `Authenticator`には以下の４つの種類がある.



1. noop: 全ての認証を通す.
```
{
    "id": "some-id",
    "upstream": {
        "url": "http://my-backend-service"
    },
    "match": {
        "url": "http://my-app/some-route",
        "methods": [
            "GET"
        ]
    },
    "authenticators": [{
        "handler": "noop"
    }]
}
```

2. anonymous: Authorization header が設定されているかをチェックする.
もしくは`AUTHENTICATOR_ANONYMOUS_USERNAME` をサブジェクト名として設定する.

```
{
    "id": "some-id",
    "upstream": {
        "url": "http://my-backend-service"
    },
    "match": {
        "url": "http://my-app/some-route",
        "methods": [
            "GET"
        ]
    },
    "authenticators": [{
        "handler": "anonymous"
    }],
    /* ... */
}
```

3. oauth2_client_credentials
username, password を使用して認証を実施.
環境変数に`AUTHENTICATOR_OAUTH2_CLIENT_CREDENTIALS_TOKEN_URL` を指定する必要がある. (like `export AUTHENTICATOR_OAUTH2_CLIENT_CREDENTIALS_TOKEN_URL=https://my-oauth2-server/oauth2/token`)

provider には, scope: `required_scope`が付随してリクエストされる.

```
{
    "id": "some-id",
    "upstream": {
        "url": "http://my-backend-service"
    },
    "match": {
        "url": "http://my-app/some-route",
        "methods": [
            "GET"
        ]
    },
    "authenticators": [{
        "handler": "oauth2_client_credentials",
        "config": {
            "required_scope": ["scope-a", "scope-b"]
        }
    }],
    /* ... */
}
```


4. oauth2_introspection
Bear Token を使用して,トークンの有効かどうかを確認する.
`oauth2_introspection` を使用する場合は, 以下４つの環境変数が必須になる.

Required:
  1. AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_ID: client の`clientid`
  2. AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_SECRET: client の`client secret`.
  3. AUTHENTICATOR_OAUTH2_INTROSPECTION_TOKEN_URL: token url.
  4. AUTHENTICATOR_OAUTH2_INTROSPECTION_INTROSPECT_URL: introspection url.

Option:
  1. AUTHENTICATOR_OAUTH2_INTROSPECTION_SCOPE: introspectionで必要な特定のスコープを設定.
  
```
{
    "id": "some-id",
    "upstream": {
        "url": "http://my-backend-service"
    },
    "match": {
        "url": "http://my-app/some-route",
        "methods": [
            "GET"
        ]
    },
    "authenticators": [{
        "handler": "oauth2_introspection",
        "config": {
            "required_scope": ["scope-a", "scope-b"]
        }
    }],
    /* ... */
}
```


#### パラメータ
* handler: `authenticator`の名前
* config: オプション.



### Authorizers
authenticator から返された`subject`を要求されたアクションを実行できるかチェックする.
以下のhandler を使用することができる.

1. allow: 全ての要求を通す.
```
    "authorizer": {
        "handler": "allow"
    }
```
2. deny: 全てのリクエストを拒否する.
```
    "authorizer": {
        "handler": "deny"
    }

```

3. keto_warden: keto api を使用して, ACLを実行.
`AUTHORIZER_KETO_WARDEN_KETO_URL` を設定する必要がある.
仮に設定されていなかった場合は, 無効になる.

２つにconfiguration が設定可能.
* required_action
* required_resource

```
"authorizer": {
    "handler": "keto_warden",
    "config": {
        "required_action": "..."
        "required_resource": "..."
    }
}
```

以下のように可変のパスにも対応可能.
```
"match": {
    "url": "http://my-app/api/users/<[0-9]+>/<[a-zA-Z]+>",
    "methods": ["GET"]
},
```

$1, $2 にそれぞれの値が入る.
```
"config": {
    "required_action": "my:action:$1",
    "required_resource": "my:resource:$2:foo:$1"
}
```


```
{
    "id": "some-id",
    "upstream": {
        "url": "http://my-backend-service"
    },
    "match": {
        "url": "http://my-app/api/users/<[0-9]+>/<[a-zA-Z]+>",
        "methods": [
            "GET"
        ]
    },
    "authenticators": [/* ... */],
    "authorizer": {
        "handler": "keto_warden",
        "config": {
            "required_action": "my:action:$1",
            "required_resource": "my:resource:$2:foo:$1"
        }
    }
    /* ... */
}
```


