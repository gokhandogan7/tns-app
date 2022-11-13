import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import SearchBox from "../../components/ui/SearchBox";
import { useNavigate } from "react-router-dom";
import "./articlepage.css";
import useDebounce from "../../hooks/useDebounce";
import Item from "../../components/ui/Item";

export const ArticlePage = () => {
  const navigate = useNavigate();

  const navigateToAuthors = () => {
    // 👇️ navigate to /contacts
    navigate("/authors");
  };
  const navigateToContents = () => {
    // 👇️ navigate to /
    navigate("/contents");
  };
  const navigateToHighlights = () => {
    // 👇️ navigate to /
    navigate("/highlights");
  };
  const navigateToHome = () => {
    // 👇️ navigate to /
    navigate("/");
  };

  const [limit, setLimit] = useState("");
  const [aqLimit, setAqLimit] = useState("");
  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 700);
  const [state] = useApi(services.getArticles, [], debouncedValue, aqLimit);

  //handleChange function for search
  const searchHandleChange = (event) => {
    setValue(event.target.value);
  };

  const limitHandleChange = (event) => {
    const re = /^[0-9\b]+$/;
    if (event.target.value === "" || re.test(event.target.value)) {
      setLimit(event.target.value);
    }
  };4454

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>List of Articles</h1>
      <SearchBox handleChange={searchHandleChange} value={value} />
      <div data-cy="limit-container" style={{ width: 200, display: "table", margin:20}}>
        <div style={{ display: "table-cell", width: "100%" }}>
          <input
          data-cy="limit-searchbox"
            value={limit}
            onChange={limitHandleChange}
            style={{ width: "100%" }}
            type="number"
            min={0}
          />
        </div>
        <input
        data-cy="limit-button"
          onClick={() => setAqLimit(limit)}
          style={{ width: "100%", marginLeft: 10 }}
          type="button"
          value="Bring"
        />
      </div>
      <button data-cy="author-button" className="buttonA buttonAA" onClick={navigateToAuthors}>
        Author
      </button>
      <button data-cy="content-button" className="button button3" onClick={navigateToContents}>
        Content
      </button>
      <button data-cy="highlight-button" className="buttonH buttonHH" onClick={navigateToHighlights}>
        Highlight
      </button>
      <button data-cy="gohome-button" className="buttonG buttonGG" onClick={navigateToHome}>
        Go Back To Home
      </button>
      <Item items={state} search={debouncedValue} />
    </div>
  );
};
