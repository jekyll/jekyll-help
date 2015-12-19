# [loop &quot;Regenerating: 0 files at ...&quot;](https://github.com/jekyll/jekyll-help/issues/49)

> state: **closed** opened by: **** on: ****

This problem appears with upgrade to Jekyll 2.
Using

    bundle exec jekyll serve --watch

When some file modified jekyll regenerating one, and then keep trying to regenerate. Then, when something new was modified, it&#x27;s regenerate the new file, and continues &quot;0 files&quot;.

    Regenerating: 1 files at 2014-05-20 16:34:32 ...done.
    Regenerating: 0 files at 2014-05-20 16:34:33 ...done.
    Regenerating: 0 files at 2014-05-20 16:34:34 ...done.

...

    Regenerating: 1 files at 2014-05-20 16:34:39 ...done.
    Regenerating: 0 files at 2014-05-20 16:34:39 ...done.
    Regenerating: 0 files at 2014-05-20 16:34:40 ...done.

Using

    bundle exec jekyll serve --watch --force_polling

solve this problem, but with this option regeneration event does not always work. 


### Comments

---
> from: [**ValeryVS**](https://github.com/jekyll/jekyll-help/issues/49#issuecomment-43623597) on: **Tuesday, May 20, 2014**

This musb in jekyll, not in jekyll-help.
Close. Moved.
