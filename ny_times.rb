require 'rest-client'
require 'json'

class Execution
  attr_accessor(:request_url, :response, :header, :raw_response)

  def initialize(year)
    api_key = "j5SaghqWK6jXAdHBXFpTSkRcdUWf5ZqJ"
    month = (1..12).to_a.sample.to_s
    @request_url = "https://api.nytimes.com/svc/archive/v1/#{year}/#{month}.json?&api-key=#{api_key}"
    @header = {"content-type" => "charset=utf-8"}
  end

  def find_article
    begin
      @raw_response = RestClient::Request.execute(method: :get, url: request_url, timeout: 10, headers: header)
    rescue RestClient::Forbidden
      @response = "Please Enter a valid year"
      output_hash = {"Error" => @response}
    else
      @response = JSON.parse(@raw_response)['response']['docs'].sample
      output_hash = {"Abstract" => @response['abstract']}
    end
    JSON.generate(output_hash)
  end
end

execution = Execution.new(ARGV[0])
print execution.find_article