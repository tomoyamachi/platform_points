# coding: utf-8
require_relative './response'
require 'rest-client'
require 'json-schema'

if ENV['cucumber_api_verbose'] == 'true'
  RestClient.log = 'stdout'
end

$cache = {}

Given(/^I send and accept JSON$/) do
  steps %Q{
    Given I send "application/json" and accept JSON
  }
end

Given(/^I send "(.*?)" and accept JSON$/) do |content_type|
  @headers = {
      :Accept => 'application/json',
      :'Content-Type' => %/#{content_type}/
  }
end

When(/^I set JSON request body to '(.*?)'$/) do |body|
  @body = JSON.parse body
end

When(/^I set form request body to:$/) do |params|
  @body = {}
  params.rows_hash.each do |key, value|
    p_value = value
    @grabbed.each { |k, v| p_value = v if value == %/{#{k}}/ } unless @grabbed.nil?
    p_value = File.new %-#{Dir.pwd}/#{p_value.sub 'file://', ''}- if %/#{p_value}/.start_with? "file://"
    @body[%/#{key}/] = p_value
  end
end

When(/^I set request body from "(.*?).(yml|json)"$/) do |filename, extension|
  path = %-#{Dir.pwd}/#{filename}.#{extension}-
  if File.file? path
    case extension
      when 'yml'
        @body = YAML.load File.open(path)
      when 'json'
        @body = JSON.parse File.read(path)
      else
        raise %/Unsupported file type: '#{path}'/
    end
  else
    raise %/File not found: '#{path}'/
  end
end


When(/^I grab "(.*?)" as "(.*?)"$/) do |json_path, place_holder|
  if @response.nil?
    raise 'No response found, a request need to be made first before you can grab response'
  end

  @grabbed = {} if @grabbed.nil?
  @grabbed[%/#{place_holder}/] = @response.get json_path
end

Then(/^the response status should be "(\d+)"$/) do |status_code|
  raise %/Expect #{status_code} but was #{@response.code}/ if @response.code != status_code.to_i
end

Then(/^the JSON response should follow "(.*?)"$/) do |schema|
  file_path = %-#{Dir.pwd}/#{schema}-
  if File.file? file_path
    begin
      JSON::Validator.validate!(file_path, @response.to_s)
    rescue JSON::Schema::ValidationError => e
      raise JSON::Schema::ValidationError.new(%/#{$!.message}\n#{@response.to_json_s}/,
                                              $!.fragments, $!.failed_attribute, $!.schema)
    end
  else
    puts %/WARNING: missing schema '#{file_path}'/
    pending
  end
end

Then(/^the JSON response root should be (object|array)$/) do |type|
  steps %Q{
    Then the JSON response should have required key "$" of type #{type}
  }
end

Then(/^the JSON response should have key "([^\"]*)"$/) do |json_path|
  steps %Q{
    Then the JSON response should have required key "#{json_path}" of type any
  }
end

Then(/^the JSON response should have (required|optional) key "(.*?)" of type \
(numeric|string|array|boolean|numeric_string|object|array|any)( or null)?$/) do |optionality, json_path, type, null_allowed|
  next if optionality == 'optional' and not @response.has(json_path)  # if optional and no such key then skip
  if 'any' == type
    @response.get json_path
  elsif null_allowed.nil?
    @response.get_as_type json_path, type
  else
    @response.get_as_type_or_null json_path, type
  end
end

# Bind grabbed values into placeholders in given URL
# Ex: http://example.com?id={id} with {id => 1} becomes http://example.com?id=1
# @param url [String] parameterized URL with placeholders
# @return [String] binded URL or original URL if no placeholders
def resolve url
  unless @grabbed.nil?
    @grabbed.each { |key, value| url = url.gsub /\{#{key}\}/, %/#{value}/ }
  end
  url
end


When(/^I send a (GET|POST|PATCH|PUT|DELETE) request to "(.*?)" with:$/) do |method, url, params|
  unless params.hashes.empty?
    query = params.hashes.first.map{|key, value| %/#{key}=#{value}/}.join("&")
    if url.include?('?')
      url = url+"&"+query
    else
      url = url+"?"+query
    end
  end

  steps %Q{
      When I send a #{method} request to "#{url}"
    }
end


When(/^I send a (GET|POST|PATCH|PUT|DELETE) request to "(.*?)"$/) do |method, url|
  # ログインしてセッションが作成されていれば、{account_id}に自分のIDを入れる
  if @account_id.nil?
    resolveUrl = URI.encode resolve(url)
  else
    resolveUrl = URI.encode resolve(url.sub(/{account_id}/,@account_id.to_s))
  end
  request_url = $BASEURL+resolveUrl

  @headers = {} if @headers.nil?
  begin
    case method
      when 'GET'
        response = RestClient.get request_url, @headers
      when 'POST'
        response = RestClient.post request_url, @body, @headers
      when 'PATCH'
        response = RestClient.patch request_url, @body, @headers
      when 'PUT'
        response = RestClient.put request_url, @body, @headers
      else
        response = RestClient.delete request_url, @headers
    end
  rescue RestClient::Exception => e
    response = e.response
  end

  @response = CucumberApi::Response.create response
  @headers = nil
  @body = nil
  @grabbed = nil
  $cache[%/#{request_url}/] = @response if 'GET' == %/#{method}/
end

Then(/^"(.*?)" should be equal "(.*?)"$/) do |key, value|
  if @grabbed.nil?
    raise %/Undefined key: '#{key}'/
  else
    # 型比較でboolean型などがあるので、jsonの値を文字列に変換して比較
    raise %/Expect #{value} but was #{@grabbed[key]}/ if @grabbed[key].to_s != value
  end
end

# セッションを保持した状態での接続
Given /^I am logged in as:$/ do |params|
  # session_idとして保存されるIDをランダムに生成
  if $current_session.nil?
    $current_session = (0...8).map { (65 + rand(26)).chr }.join
  end

  # 指定されたparameterでログインをテスト
  # When I am logged in as:
  #   | mail     | test@mail.com |
  #   | password | samplepass    |
  @body = {}
  params.rows_hash.each do |key, value|
    p_value = value
    @grabbed.each { |k, v| p_value = v if value == %/{#{k}}/ } unless @grabbed.nil?
    p_value = File.new %-#{Dir.pwd}/#{p_value.sub 'file://', ''}- if %/#{p_value}/.start_with? "file://"
    @body[%/#{key}/] = p_value
  end

  @headers = {
              :Accept => 'application/json',
              :'Content-Type' => 'application/json',
              :Cookie => $SESSION_NAME+'='+$current_session,
  }

  # POSTでログインJSONレスポンスをパースしてアカウントIDを指定する
  response = RestClient.post $BASEURL+$LOGIN_ENDPOINT, @body, @headers
  @response = CucumberApi::Response.create response
  parsed = JSON.parse @response

  # 返却データから、本人のアカウントIDを保存する
  @account_id = parsed['id']
  @body = nil

  # 次回のリクエストでも、同じセッションで接続する(ログイン状態とみなされる)
  @headers = {
              :Accept => 'application/json',
              :'Content-Type' => 'application/json',
              :Cookie => $SESSION_NAME+'='+$current_session,
  }
end
