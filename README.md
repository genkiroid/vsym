# Vsym

Vsym is a Go tool for checking TLS certificates that may not be trusted by Chrome.

**Because this tool is based solely on [published information](https://security.googleblog.com/2017/09/chromes-plan-to-distrust-symantec.html), it will not display exactly the same result as Chrome's decision. Please use it at your own risk.**

## Installation

Precompiled binaries for released versions are available in the [releases](https://github.com/genkiroid/vsym/releases) page.

## Usage

Give domain names as arguments.

```sh
$ vsym securitycenter.rapid-ssl.jp seal.websecurity.norton.com
The SSL certificate used on https://securitycenter.rapid-ssl.jp will be distrusted in Chrome v66.
The SSL certificate used on https://seal.websecurity.norton.com will be distrusted in Chrome v70.
```

## License

[MIT](https://github.com/genkiroid/vsym/blob/master/LICENSE)

## Author

[genkiroid](https://github.com/genkiroid)
