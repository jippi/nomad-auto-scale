# nomad-auto-scale

This tool is in super alpha, its not even close to be functional yet :)

It's an attempt for me to port the PHP nomad auto-scaler I use internally at bownty.com.

The idea is you specify your job configuration in either a config file (see `example-configs/` directory), or let
the job spec be part of the nomad job file meta stanza.

For now I'm focusing on the scale rules in a HCL file, once that works, it should be fairly trivial to add Nomad Job backend.

My first priority is to implement scale up/down based on RabbitMQ queue size and utilization. Secondly a HTTP backend would be
nice, as it would allow this tool to hit any HTTP API and get directions on scale up/down, leaving you as user of the tool
to implement your own complex logic in whatever language you have a strongest preference for.

Other backends could be Graphite, DataDog, NewRelic or Prometheus, though external contributions probably would be needed
for that.

Feel free to hit me up on Gitter (both in the Nomad and Hashi-UI room).