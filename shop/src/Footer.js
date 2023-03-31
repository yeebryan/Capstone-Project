import React from 'react';

import "./App.css";

// footer
export default function footer1() {
    
    return (
      <footer className="bg-primary-500 text-white p-4">
        <div className="container mx-auto">
          <div className="flex flex-wrap items-center justify-between">
            <table className="footer-table">
              <tr className="footer-stuff">
                <td className="footer-td">
                <p className="text-sm">© 2023 Your Company. All rights reserved.</p>
                </td>
                <td className="footer-tdd">
                  <a href="/" className="text-black hover:text-primary-100">
                    Terms of Service
                  </a> 
                </td>
                <td className="footer-tdd">
                  <a href="/" className="text-black hover:text-primary-100">
                    Privacy Policy
                  </a>
                </td>
                <td className="footer-tdd">
                  <a href="/" className="text-black hover:text-primary-100">
                    Contact Us
                  </a>
                </td>
              </tr>
            </table>
          </div>
        </div>
      </footer>
    );
  };

  