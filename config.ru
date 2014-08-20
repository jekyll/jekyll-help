
require 'bundler'
Bundler.require(:default, :production)

run Rack::Jekyll.new
