import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import SearchBox from "../../components/ui/SearchBox";
import { useNavigate } from "react-router-dom";
import "./contentpage.css"
import useDebounce from "../../hooks/useDebounce";

export const ContentPage = () => {
  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 700);
  const [contents] = useApi(services.getAllContents, [], debouncedValue);
  const handleChange = (event) => {
    setValue(event.target.value);
  };

  const navigate = useNavigate();

  const navigateToHome = () => {
    // 👇️ navigate to /contacts
    navigate("/fullarticles");
  };
  console.log(contents)
  return (
    <>
      <div style={{ textAlign: "center" }}>
        <h1 style={{ textAlign: "center" }}>List of Content</h1>
        <SearchBox handleChange={handleChange} value={value} />
        <button className="buttonC buttonCC" onClick={navigateToHome}>Go Back Articles Page</button>
        {contents.map((content, index) => (
          <div data-cy="content" className="container" key={index}>
            <div className="itemContainer">
              <p data-cy="text" className="text">
                {content.Text}
              </p>
              <img src={content.Image} style={{width:400, height:300}}/>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
