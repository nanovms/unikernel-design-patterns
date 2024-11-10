This example shows you how to create a background job worker using
unikernels.

Unikernels are a fantastic resource to use for background workers as
they abstract away a lot of orchestration/scheduling that one would
otherwise need to explicitly handle.

This examples show a very simple background worker abstraction.

The server is the queue server.

The worker is the job processor.

Ideas to Extend:
---------------------

* Make the server/worker use dns.
* Make the worker work across various clouds.
* Make the worker use different languages with different worker images.
