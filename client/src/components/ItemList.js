import React, { useState } from "react";
import { useApi } from "../hooks/useApi";
import { services } from "../services";
import Item from "./ui/Item";
import "./itemlist.css";

function ItemList() {
  const { state } = useApi(services.getArticles, [], false);
  const [query, setQuery] = useState("");

  return (
    <div style={{ textAlign: "center" }}>
      <h1 style={{ textAlign: "center" }}>List of Data</h1>
      <input
       className="inputBox"
        placeholder="Search Items"
        onChange={(event) => setQuery(event.target.value)}
      />
      
      <Item items={state} search={query} />
    </div>
  );
}

export default ItemList;
