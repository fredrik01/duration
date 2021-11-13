# Duration

CLI tool to parse ISO 8601 durations.

Usage:
```shell
echo "PT3H7M38S" | duration parse
# 03:07:38
```

## Install / Update

MacOS

```
curl -L https://github.com/fredrik01/duration/releases/latest/download/duration_darwin_amd64.tar.gz -o duration.tar.gz
mkdir /tmp/duration
tar -xvf duration.tar.gz -C /tmp/duration
mv /tmp/duration/duration /usr/local/bin/duration
rm duration.tar.gz
rm -r /tmp/duration
```
