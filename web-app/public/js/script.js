const url = new URLSearchParams(location.search);
const token = url.get("token");

document
  .getElementById("logout")
  ?.setAttribute("href", "/logout?token=" + token);
