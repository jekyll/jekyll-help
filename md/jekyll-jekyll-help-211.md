# [Viewing Jekyll Locally From Vagrant VM](https://github.com/jekyll/jekyll-help/issues/211)

> state: **open** opened by: **** on: ****

Hi

I have installed jekyll on a vagrant vm instance and all is good from within side the vagrant box

I can create a new jekyll site and serve it

Configuration file: /var/www/awesome/demo/_config.yml
            Source: /var/www/awesome/demo
       Destination: /var/www/awesome/demo/_site
      Generating...
                    done.
 Auto-regeneration: enabled for &#x27;/var/www/awesome/demo&#x27;
Configuration file: /var/www/awesome/demo/_config.yml
    Server address: http://127.0.0.1:4000/
  Server running... press ctrl-c to stop.

However when i try to view this on my local browser i just get 

The connection was reset

i have the ports forwarded in the vm 

config.vm.network :forwarded_port, host: 4000, guest: 4000

So do not understand why i cannot view the site locally

please can someone advise.

Regards

Mark

### Comments

