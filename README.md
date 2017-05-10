# notify

notify is Stdin to slack.

# Install

go get github.com/decafe09/notify

# Usage

set `SLACK_WEBHOOK_URL` environment variable.

```
export SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX
```

notify to general.

```
ls | notify -t "Look this!" -l good -c general
```

notify to general with warning.

```
ls | notify -t "Look this!" -l warning -c general
```

notify to general with danger.

```
ls | notify -t "Look this!" -l danger -c general
```

notify to general without title.

```
ls | notify -l good -c general
```

notify to general minimum.

```
ls | notify -c general
```

notify to general without title.

```
ls | notify -l good -c general
```


# Help

```
Usage of notify:
command | notify -t [title] -l [level] -c [channel]
  -c string
      notify channel
  -h	show help
  -l string
      set level [good|warning|danger] (default "#dddddd")
  -t string
      set title
  -v  show version
```


