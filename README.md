Trying to figure out why user.Current() is returning nil under a new release of golang and appengine: https://groups.google.com/d/msg/google-appengine-go/etliynWyeVI/2L7ncj_aCQAJ

Here's the broken commit that prompted this: @jimhopp/frederic@6407202a75a172e64b03f7d0d22b861d92c62c45

(Turns out I was abusing request headers in my test setup. Not sure why this worked pre-1.9.68, but playing nicely with the request headers solved the problem.)
