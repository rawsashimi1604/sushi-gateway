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
  return HttpRequest.get("/gateway", {
    withCredentials: true,
  });
}

function getGatewayConfig() {
  return HttpRequest.get("/gateway/config", {
    withCredentials: true,
  });
}

export default {
  login,
  logout,
  getGatewayData,
  getGatewayConfig,
};
