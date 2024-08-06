import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import * as yup from "yup";

import AdminAuth from "../../api/services/admin/AdminAuth";
import TextInput from "../../components/input/TextInput";
import LoadingSpinner from "../../components/utils/LoadingSpinner";
import { delay } from "../../utils/delay";
import Logo from "../../components/sushi-gateway/Logo";

// Define validation schema using yup for login
const loginValidationSchema = yup.object().shape({
  username: yup.string().required("Username is required."),
  password: yup.string().required("Password is required."),
});

type LoginState = {
  status: boolean;
  message: string;
};

function LoginForm() {
  const navigate = useNavigate();
  // TODO: change to auth check route...
  const [loginData, setLoginData] = useState({
    username: "",
    password: "",
  });

  const [validationErrors, setValidationErrors] = useState<{
    [key: string]: string;
  }>({});

  const [loginState, setLoginState] = useState<LoginState | null>(null);

  const handleChange = (name: string, value: string) => {
    setLoginData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoginState(null);
    try {
      await loginValidationSchema.validate(loginData, { abortEarly: false });
      setValidationErrors({});
    } catch (err) {
      if (err instanceof yup.ValidationError) {
        const errorObj: { [key: string]: string } = {};
        for (let error of err.inner) {
          errorObj[error.path as string] = error.message;
        }
        setValidationErrors(errorObj);
      }
    }

    try {
      const loginRes = await AdminAuth.login(
        loginData.username,
        loginData.password
      );
      localStorage.setItem("jwt-token", JSON.stringify(loginRes.data));
      setLoginState({
        status: true,
        message: "successfully logged in, redirecting...",
      });
      await delay(2000);
      navigate("/");
    } catch (err) {
      setLoginState({
        status: false,
        message: "invalid credentials, please try again.",
      });
    }
  };

  return (
    <div className="font-sans flex  gap-4 items-center justify-center w-screen h-screen">
      <div className="flex flex-col gap-4">
        <Logo />

        <form className="flex flex-col gap-3" onSubmit={handleLogin}>
          {/* Username */}
          <div className="flex flex-col justify-center gap-1 text-sm">
            <label
              htmlFor="username"
              className="tracking-wide flex flex-row items-center justify-start gap-3"
            >
              <span className="mb-1 ">username</span>
            </label>
            <div className="">
              <TextInput
                id="username"
                name="username"
                value={loginData.username}
                onChange={(e) => handleChange("username", e.target.value)}
                error={validationErrors.username}
              />
            </div>
          </div>

          {/* Password */}
          <div className="flex flex-col justify-center gap-1 text-sm">
            <label
              htmlFor="password"
              className="tracking-wide flex flex-row items-center justify-start gap-3"
            >
              <span className="mb-1 ">password</span>
            </label>
            <div className="">
              <TextInput
                type="password"
                id="password"
                name="password"
                value={loginData.password}
                onChange={(e) => handleChange("password", e.target.value)}
                error={validationErrors.password}
              />
            </div>
          </div>

          <button
            type="submit"
            className="w-[80px] mt-2 text-xs py-1.5 px-2 pb-2 text-white bg-blue-500 shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg"
          >
            <span>login</span>
          </button>

          {loginState && (
            <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
              {loginState.status ? (
                <React.Fragment>
                  <span className="text-green-500">{loginState?.message}</span>
                  <LoadingSpinner size={12} />
                </React.Fragment>
              ) : (
                <span className="text-red-500">{loginState.message}</span>
              )}
            </div>
          )}
        </form>
      </div>
    </div>
  );
}

export default LoginForm;
