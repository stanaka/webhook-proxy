webhook-proxy [![Build Status](https://travis-ci.org/stanaka/webhook-proxy.svg?branch=master)](https://travis-ci.org/stanaka/webhook-proxy)
=============

written in Go-lang

webhook-proxy receives webhooks from any services and delegate them to regsistered end points.

Configuration
=============

Usage
=====

```
webhook-proxy -conf conf.yml
```

Use case
========

* Webhook-proxy receive webhook from GitHub or other services.
* Then webhook-proxy notify it to internal IRC.
* Messages can be customized.

Security
========

API Key is embedded into url path.
