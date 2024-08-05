import HttpRequest from "../../requests/HttpRequest";

// function login(username: string, password: string) {
//   return HttpRequest.post<LoginResponse>(`/login`, {
//     username,
//     password,
//   });
// }

function getGatewayData() {
  return HttpRequest.get("/");
}

export default {
  getGatewayData,
};
