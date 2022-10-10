import { services } from "../../services";
import { useApi } from "../../hooks/useApi";
import "./home.css";

export const HomePage = () => {
  const {state} = useApi(services.getArticles, [])
  
  
  return (
    <div className="homePageWrapper">
      <div className="contentContainer">
        <h1 className="header">Welcome to T&S</h1>
        <p className="motto">Let's make it Possible!!!</p>
      </div>

      <div className="linkContainer">
        <a className="link" href="/details">
          Details
        </a>
      </div>
    </div>
  );
};
