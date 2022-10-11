import React, { useState } from "react";
import { useApi } from "../hooks/useApi";
import { services } from "../services";
import Item from "./ui/Item";
import "./itemlist.css";
import useDebounce from "../hooks/useDebounce";
import SearchBox from "./ui/SearchBox";


function ItemList() {
  const { state } = useApi(services.getArticles, [], false);
  const [value, setValue] = useState('')
  const debouncedValue = useDebounce(value, 1500)

  const handleChange = (event) => {
    setValue(event.target.value)
  }

  return (
    <div style={{ textAlign: "center" }}>
      <h1 style={{ textAlign: "center" }}>List of Data</h1>
      <SearchBox handleChange={handleChange} value={value}/>
      <Item items={state} search={debouncedValue} />
    </div>
  );
}



export default ItemList;
