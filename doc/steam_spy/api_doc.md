
  This is an API for Steam Spy. It accepts requests in GET string and returns data in JSON arrays.

  Allowed poll rate - 4 requests per second.

  ## Examples: ##
   
  * steamspy.com/api.php?request=appdetails&appid=730 - returns data for Counter-Strike: Global Offensive
  * steamspy.com/api.php?request=top100in2weeks - return Top 100 apps by players in the last two weeks


  ## Common parameters: ##
 
  * request - code for API request call.
  * appid - Application ID (a number).


  ## Accepted requests: ##
  
  ### appdetails ###

  Returns details for the specific application. Requires *appid* parameter.  

  ### genre ###

  Returns games in this particular genre. Requires *genre* parameter and works like this:
  
  * steamspy.com/api.php?request=genre&genre=Early+Access

  ### tag ###

  Returns games in this particular tag. Requires *tag* parameter and works like this:
  
  * steamspy.com/api.php?request=tag&tag=Bowling

  ### top100in2weeks ###

  Returns Top 100 games by players in the last two weeks.

  ### top100forever ###

  Returns Top 100 games by players since March 2009.

  ### top100owned ###

  Returns Top 100 games by owners.

  ### all ###

  Returns all games with owners data sorted by owners.


  ## Return format for an app: ##

  * appid - Steam Application ID. If it's 999999, then data for this application is hidden on developer's request, sorry.
  * name - the game's name
  * developer - comma separated list of the developers of the game
  * publisher - comma separated list of the publishers of the game
  * score_rank - score rank of the game based on user reviews
  * positive - number of positive user reviews
  * negative - number of negative user reviews
  * userscore - user score of the game
  * owners - owners of this application on Steam. **Beware of free weekends!**
  * owners_variance - variance in owners. The real number of owners lies somewhere on owners +/- owners_variance range.   
  * players_forever - people that have played this game since March 2009.
  * players_forever_variance - variance for total players.
  * players_2weeks - people that have played this game in the last 2 weeks.
  * players_2weeks_variance - variance for the number of players in the last two weeks. 
  * average_forever - average playtime since March 2009. In minutes.
  * average_2weeks - average playtime in the last two weeks. In minutes.
  * median_forever - median playtime since March 2009. In minutes.
  * median_2weeks - median playtime in the last two weeks. In minutes.
  * price - US price in cents.

  * ccu - peak CCU yesterday (only returned if an individual app is requested).
  * tags - the game's tags with votes in JSON array (only returned if an individual app is requested).


  ## Questions? ##

  Contact me by e-mail: *sergey at galyonkin dot com*.

  