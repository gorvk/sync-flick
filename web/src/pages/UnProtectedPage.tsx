import { Link } from "react-router-dom";

export const UnProtectedPage = () => (
  <>
    <div className="text-2xl">This is a unprotected page</div>
    <Link className="underline decoration-blue-500 text-blue-500" to="/protected">
      click here to navigate to protected page with api data
    </Link>
  </>
);
