# Golab.io 2024-11-12 - Let's go asynchronous

See the presentation in the `slides` directory.
The original presentation version is in <a href="https://docs.google.com/presentation/d/1UgxxIQvfHi9v9GiaeQWMfmN1FwowdShJR9nVyEKLnAo/edit?usp=sharing">Google Slides</a>.

Code examples:
- Synchronous code:
  - Invoice manually: `sync/http`
  - Blocked by printer: `sync/http-printer`
- Asynchronous code:
  - Invoice asynchronously over HTTP with mem leak: `async/http`
  - Invoice asynchronously over AMQP: `async/rabbit`
  - Invoice asynchronously over PGQ: `async/pgq`

In each example directory, run `make build` to build and `make run` to run the example.