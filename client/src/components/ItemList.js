import React from "react";
import { useApi } from "../hooks/useApi";
import { services } from "../services";
import Item from "./ui/Item";

function ItemList() {
  const { state } = useApi(services.getArticles, [], false);
  console.log(state);
  return (
    <div>
      <>
        <h1 style={{ textAlign: "center" }}>List of Data</h1>
        <Item items={state} />
      </>
    </div>
  );
}

export default ItemList;
