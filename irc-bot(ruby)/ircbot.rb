require "cinch"
require 'cinch'
require 'open-uri'
require 'nokogiri'
require 'cgi'

bot = Cinch::Bot.new do
  configure do |c|
	  c.nick="osdc-bot"
    c.server   = "irc.freenode.net"
    c.channels = ["#jiit-lug"]
  end
 helpers do

  def urban_dict(query)
      url = "http://www.urbandictionary.com/define.php?term=#{CGI.escape(query)}"
      CGI.unescape_html Nokogiri::HTML(open(url)).css("div.meaning").first.text.gsub(/\s+/, ' ').strip rescue nil
    end

  def google(query)
      url = "http://www.google.com/search?q=#{CGI.escape(query)}"
      res = Nokogiri.parse(open(url).read).at("h3.r")

      title = res.text
      link = res.at('a')[:href]
      desc = res.at("./following::div").children.first.text
    rescue
      "No results found"
    else
      CGI.unescape_html "#{title} - #{desc} (#{link})"
end

end

  	on :message, "!hello" do |m|
	  	m.reply "Hello #{m.user.nick},welcome to OSDC channel on iRc"
	 end

  	on :message, "!whoareyou" do |m|
	  	m.reply "#{m.user.nick},I am time,space and everything."
	 end

	on :message, /^!mean (.+)/ do |m, term|
    	m.reply(urban_dict(term) || "No results found", true)
	 end

	on :message, /^!google (.+)/ do |m, query|
    m.reply google(query)
end
end

bot.start