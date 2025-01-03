import { useEffect } from "react";
import { useDispatch } from "react-redux";
import { IdToken, useAuth0 } from "@auth0/auth0-react";
import { slice as userSlice } from "../state/slices/user";

export function Auth() {
  const { loginWithRedirect, getIdTokenClaims, isAuthenticated } = useAuth0();
  const dispatch = useDispatch();

  useEffect(() => {
    authenticateUser();
  }, [isAuthenticated]);

  async function authenticateUser() {
    if (!isAuthenticated) {
      await loginWithRedirect();
    }
    const claim = await getIdTokenClaims();
    const idToken = claim?.__raw;
    if (idToken) {
      localStorage.setItem("id_token", idToken);
      dispatch(userSlice.actions.setCurrentUser(claim as IdToken));
    }
  }

  return <></>;
}
