function IsUnixTimeExpired(unixTimestamp) {
  const currentTime = Math.floor(Date.now() / 1000); 

  return unixTimestamp < currentTime
}

module.exports = {
  IsUnixTimeExpired: IsUnixTimeExpired
};