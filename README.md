# Cronjob Istio to Linkerd KillProxy

## Description

When using Cronjobs with a Service Mesh Istio or Linkerd, you must make a HTTP request to the proxy's endpoint (e.g `localhost:15020/quitquitquit` on istio or `localhost:4191/shutdown` on Linkerd) to terminate the proxy when you Job as completed it's task.

This docker image impersonates the istio proxy for the shutdown phase. When disabling istio injection your app containers may still curl the localhost:15020/quitquitquit even though istio is not injected, resulting in the failure of all your cronjobs, this image fixes this issue.

## Upcoming / TODO

Add a `wget --post-data hello=shutdown http://localhost:4191/shutdown` in the quitquitquit router to be able to migrate from istio to linkerd without having to redeploy the CronJobs