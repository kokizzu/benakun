function IsUnixTimeExpired(unixTimestamp) {
  const currentUnixTime = Math.floor(Date.now() / 1000);

  return unixTimestamp < currentUnixTime;
}

module.exports = {
  IsUnixTimeExpired: IsUnixTimeExpired
};