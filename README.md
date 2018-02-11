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

Give domain list. [example](https://github.com/genkiroid/vsym/blob/master/examples) is example file.

```sh
$ vsym `cat examples`
The SSL certificate used on https://securitycenter.rapid-ssl.jp will be distrusted in Chrome v66.
The SSL certificate used on https://seal.websecurity.norton.com will be distrusted in Chrome v70.
The SSL certificate used on https://img.en25.com will be distrusted in Chrome v70.
The SSL certificate used on https://tracker.mrpfd.com will be distrusted in Chrome v70.
The SSL certificate used on https://s1701211846.t.eloqua.com will be distrusted in Chrome v70.
The SSL certificate used on https://s912704989.t.eloqua.com will be distrusted in Chrome v70.
The SSL certificate used on https://s.adroll.com will be distrusted in Chrome v70.
The SSL certificate used on https://dsum-sec.casalemedia.com will be distrusted in Chrome v70.
The SSL certificate used on https://us-u.openx.net will be distrusted in Chrome v70.
The SSL certificate used on https://ib.adnxs.com will be distrusted in Chrome v70.
```

## License

[MIT](https://github.com/genkiroid/vsym/blob/master/LICENSE)

## Author

[genkiroid](https://github.com/genkiroid)
