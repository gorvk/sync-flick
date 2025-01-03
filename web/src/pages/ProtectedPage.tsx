import { useEffect, useRef, useState } from "react";
import { Link } from "react-router-dom";
import { addData, getAllData } from "../svc/data";
import { Data } from "../models/types";
import "../styles.css";
import { useSelector } from "react-redux";
import { AppState } from "../state/store";
import { useAuth0 } from "@auth0/auth0-react";

export const ProtectedPage = () => {
  const [data, setData] = useState<Data[]>([]);
  const inputRef = useRef<HTMLInputElement>(null);
  const currentUser = useSelector((state: AppState) => state.currentUser);
  const { logout } = useAuth0();

  useEffect(() => {
    getData();
  }, []);

  const logoutUser = async () => {
    localStorage.removeItem("id_token");
    logout({ logoutParams: { returnTo: window.location.origin } });
  };

  const submitData = async () => {
    if (inputRef.current && inputRef.current.value) {
      const data: Data = {
        id: inputRef.current?.value || "",
      };
      await addData(data);
      await getData();
      inputRef.current.value = "";
    }
  };

  const getData = async () => {
    const response = await getAllData();
    setData(response);
  };

  return (
    <div className="m-4">
      <div>
        <div className="flex gap-4 items-center my-2">
          <p>
            {currentUser?.nickname} - {currentUser?.name}
          </p>
          <button className="btn" onClick={logoutUser}>
            Logout
          </button>
        </div>
        <div className="text-2xl">
          This is a protected page and cannot be accessed without login
        </div>
        <Link className="underline decoration-blue-500 text-blue-500" to="/">
          click here to navigate to unprotected page with api data
        </Link>
      </div>
      <br />
      <div>
        <div className="text-2xl">API Data:</div>
        <ul className="list-disc pt-4 pl-8">
          {data.map((d, index) => (
            <li key={index}>{d.id}</li>
          ))}
        </ul>
      </div>
      <br />
      <div>
        <div className="text-2xl">Insert Data:</div>
        <div className="flex gap-4">
          <input
            className="shadow appearance-none border rounded w-1/2 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            type="text"
            ref={inputRef}
            placeholder="random data*"
          />
          <button className="btn" type="submit" onClick={submitData}>
            Submit
          </button>
        </div>
      </div>
    </div>
  );
};
