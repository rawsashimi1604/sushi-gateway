import HttpRequest from "../../requests/HttpRequest";

function login(username: string, password: string) {
  return HttpRequest.post(
    "/login",
    {},
    {
      headers: {
        Authorization: `Basic ${btoa(`${username}:${password}`)}`,
      },
      withCredentials: true,
    }
  );
}

function logout() {
  return HttpRequest.delete("/logout", { withCredentials: true });
}

function getGatewayData() {
  return HttpRequest.get("/", {
    withCredentials: true,
  });
}

export default {
  login,
  logout,
  getGatewayData,
};
