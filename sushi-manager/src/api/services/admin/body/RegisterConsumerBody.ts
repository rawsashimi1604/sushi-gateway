export type RegisterConsumerBody = {
  username: string;
  services: number[];
  enableJwtAuth: boolean;
  jwtCredentialsName: string;
};
