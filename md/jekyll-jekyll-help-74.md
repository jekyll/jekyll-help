# [Multiple config files overriden or merged?](https://github.com/jekyll/jekyll-help/issues/74)

> state: **closed** opened by: **** on: ****

If I run

&#x60;&#x60;&#x60;
jekyll build --config _config_one.yml,_config_two.yml
&#x60;&#x60;&#x60;

does the second config fully override the first or are both merged? If merged, are they deep merged?

### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/74#issuecomment-46129793) on: **Sunday, June 15, 2014**

They are deep merged: https://github.com/jekyll/jekyll/blob/master/lib/jekyll/configuration.rb#L168
