import { Routes, Route, Navigate } from "react-router-dom";
import { useSelector } from "react-redux";
import { AppState } from "../../state/store";
import { useAuth0 } from "@auth0/auth0-react";
import { Auth } from "../../pages/Auth";
import { UnProtectedPage } from "../../pages/UnProtectedPage";
import { ProtectedPage } from "../../pages/ProtectedPage";

export const AppRoutes = () => {
  const currentUser = useSelector((state: AppState) => state.currentUser);
  const { isLoading } = useAuth0();

  if(isLoading) {
    return <h3 className="text-white">LOADING...</h3>;
  }

  return (
    <Routes>
      <Route path="/" element={<UnProtectedPage />} />
      <Route
        path="/auth"
        element={!currentUser ? <Auth /> : <Navigate to="/protected" />}
      />
      <Route
        path="/protected"
        element={currentUser ? <ProtectedPage /> : <Navigate to="/auth" />}
      />
    </Routes>
  );
};
