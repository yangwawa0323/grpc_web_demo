
import client from '../../../../js/user/v1/user_service_grpc_web_pb'

export const newUserService = () => {
  return new client.UserSearchServiceClient("http://localhost:8199")
}