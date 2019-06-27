# kafka-notifier (stage of this project: proposal)

This project is basically a proxy to abstract kafka away.

The reason for this project is because it's easy to write the producer but much harder to
do the consumer right (such as handle the offset, recovery, failures, etc correctly).
It's better to create a project to handle the theory part and let other
projects focus on their business logic, with this project other projects don't have to include
any kafka library, the only thing for them is to create http handler to handle the event they
interested in.

For the browser, you can use [sock.js](https://github.com/sockjs/sockjs-client) to connect to
this service to subscribe events.

```text
          publish topic         subscribe topics                   http request
serviceA ---------------> kafka -----------------> kafka-notifier --------------> serviceC
                            ^                            ^
           publish topic    |                            | websocket
serviceB -------------------+                         browser
```

When kafka-notifier starts it will read the yaml config file for the proxy rules.

For example the serviceA publishes topic `payment-done { id, card_token }`,
the browser and the serviceC both want to subscribe this topic, of cause we don't want to
send all the card_token back to browser, so the proxy rules is basically used to filter
the value of each message.

The rules for this example might look like:

```yaml
- targets:
  - serviceC/api/payments-events # send http request to serviceC
  topics:
  - name: payment-done
    json_filter:
    - id
    - card_token
    - customer.name # json path

- endpoint: /payments-browser # the websocket endpoint
  topics:
  - name: payment-done
    json_filter:
    - id
```

The event send from kafka-notifier will be like below. All data will be json format.

For the browser:

```json
{
    "topic": "payments-done",
    "value": {
        "id": "123"
    }
}
```

For the serviceC:

```json
{
    "topic": "payments-done",
    "value": {
        "id": "123",
        "card_token": "abc",
        "customer": {
            "name": "jack"
        }
    }
}
```

## Development

Only golang and docker are required.

### Test

```bash
go get github.com/ysmood/gokit/cmd/godev
godev test
```

### Build

```bash
godev build
```
