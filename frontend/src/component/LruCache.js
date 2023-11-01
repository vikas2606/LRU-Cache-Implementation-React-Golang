import React, { useEffect, useState } from "react";
import axios from "axios";

function LruCache() {
  const [keys, setKeys] = useState([]);

  const [getKey, setGetKey] = useState("");
  const [setKey, setSetKey] = useState("");
  const [setValue, setSetValue] = useState("");
  const [expiration, setExpiration] = useState();
  const [cachedValue, setCachedValue] = useState("");
  const [getError, setGetError] = useState(null);
  const [SetError, setSetError] = useState(null);
  const [successMessage, setSuccessMessage] = useState("");

  useEffect(() => {
    listKeys();
  }, [cachedValue,setValue,setKey]);

  const handleGet = () => {
    axios
      .get(`/get/${getKey}`)
      .then((response) => {
        setCachedValue(response.data);
        setGetError(null);
      })
      .catch((error) => {
        console.log(error);
        setCachedValue("");
        setGetError("Key not found in cache");
      });
  };

  const handleSet = (e) => {
    e.preventDefault();
    axios
      .post(`/set`, {
        key: setKey,
        value: setValue,
        expiration: expiration,
      })
      .then(() => {
        setCachedValue("");
        setSetError(null);
        setSuccessMessage("Cache value successfully added.");
      })
      .catch((error) => {
        console.log(error);
        setCachedValue("");
        setSetError(error.response.data.error);
      });
  };

  const listKeys = () => {
    axios
      .get(`/get/`)
      .then((response) => {
        setKeys(response.data);
        console.log(response)
      })
      .catch((error) => {
        console.error("Error fetching keys:", error);
      });
  };

  return (
    <div>
      <h1>LRU Cache</h1>
      <div>
        <h2>Get Value</h2>
        <div>
          <label>Key:</label>
          <input
            type="text"
            value={getKey}
            onChange={(e) => setGetKey(e.target.value)}
          />
        </div>
        <button onClick={handleGet}>Get Value</button>
        {cachedValue && <div>Cached Value: {cachedValue}</div>}
        {getError && <div className="error">{getError}</div>}
      </div>

      <form onSubmit={handleSet}>
        <h2>Set Value</h2>
        <div>
          <label>Key:</label>
          <input
            type="text"
            value={setKey}
            required
            onChange={(e) => setSetKey(e.target.value)}
          />
        </div>
        <div>
          <label>Value:</label>
          <input
            required
            type="text"
            value={setValue}
            onChange={(e) => setSetValue(e.target.value)}
          />
        </div>
        <div>
          <label>Expiration (seconds):</label>
          <input
            type="number"
            value={expiration}
            onChange={(e) => setExpiration(e.target.value)}
          />
        </div>
        <button type="submit">Set Value</button>
      </form>
      {SetError && <div className="error">{SetError}</div>}
      {successMessage && <div className="success">{successMessage}</div>}
      <h2>Keys in Cache:</h2>
      <ul>
        {keys.map((key, index) => (
          <li key={index}>{key}</li>
        ))}
      </ul>
      <button onClick={listKeys}>Refresh</button>
    </div>
  );
}

export default LruCache;
