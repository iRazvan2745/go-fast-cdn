'use client'
import { useEffect, useState } from "react";

export default function Status() {
  const [status, setStatus] = useState("Offline");
  
  const fetchStatus = async () => {
    try {
      const response = await fetch("http://localhost:8080/api");
      if(response.status === 200){
        setStatus("Online");
      } else {
        setStatus("Offline");
      }
    } catch (error) {
      setStatus("Offline");
    }
  };
  
  useEffect(() => {
    fetchStatus();
    const interval = setInterval(() => {
      fetchStatus();
    }, 3000);
    return () => clearInterval(interval);
  }, []);
  
  return (
    <div className="flex items-center gap-2">
      {status === "Online" ? (
        <div className="relative">
          <div className="h-2 w-2 rounded-full bg-green-500" />
          <div className="absolute -top-0.5 -left-0.5 h-3 w-3 animate-ping rounded-full bg-green-500 opacity-75" />
        </div>
      ) : (
        <div className="relative">
          <div className="h-2 w-2 rounded-full bg-red-500" />
        </div>
      )}
      <p className="text-sm font-medium">{status}</p>
    </div>
  )
}