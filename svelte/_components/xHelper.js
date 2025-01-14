function IsUnixTimeExpired(unixTimestamp) {
  const currentTime = new Date().getTime();

  return unixTimestamp < currentTime
}

module.exports = {
  IsUnixTimeExpired: IsUnixTimeExpired
};