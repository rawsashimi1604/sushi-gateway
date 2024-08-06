import HttpRequest from "../../requests/HttpRequest";

function login(username: string, password: string) {
  return HttpRequest.post(
    "/login",
    {},
    {
      headers: {
        Authorization: `Basic ${btoa(`${username}:${password}`)}`,
      },
    }
  );
}

function getGatewayData() {
  return HttpRequest.get("/");
}

export default {
  login,
  getGatewayData,
};
