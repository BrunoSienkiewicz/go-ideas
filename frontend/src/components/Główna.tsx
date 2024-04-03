import React from "react";
import NavBar from "./NavBar";

const HomePage = () => {
  return (
    <html lang="en">
      <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <meta charSet="utf-8" />
        <meta name="keywords" content="" />
        <meta name="description" content="" />
        <title>Główna</title>
        <link rel="stylesheet" href="Główna.css" media="screen" />
        <meta name="theme-color" content="#478ac9" />
        <meta property="og:title" content="Główna" />
        <meta property="og:description" content="" />
        <meta property="og:type" content="website" />
      </head>
      <body
        data-path-to-root="./"
        data-include-products="false"
        className="u-body u-xl-mode"
        data-lang="en"
      >
        <NavBar />
        <header className="u-clearfix u-header u-sticky u-header" id="sec-6944">
          <div className="u-clearfix u-sheet u-valign-middle u-sheet-1">
            <a
              href="#"
              className="u-image u-logo u-image-1"
              data-image-width="5000"
              data-image-height="5000"
            >
              <img
                className="u-logo-image u-logo-image-1"
                src="images/41422377-0.png"
              />
            </a>
            <div
              className="menu-collapse"
              style={{ fontSize: "1rem", letterSpacing: "0px" }}
            >
              <a
                className="u-button-style u-custom-left-right-menu-spacing u-custom-padding-bottom u-custom-top-bottom-menu-spacing u-nav-link u-text-active-palette-1-base u-text-hover-palette-2-base"
                href="#"
              >
                <svg className="u-svg-link" viewBox="0 0 24 24">
                  <use
                    xmlnsXlink="http://www.w3.org/1999/xlink"
                    xlinkHref="#menu-hamburger"
                  ></use>
                </svg>
                <svg
                  className="u-svg-content"
                  version="1.1"
                  id="menu-hamburger"
                  viewBox="0 0 16 16"
                  x="0px"
                  y="0px"
                  xmlnsXlink="http://www.w3.org/1999/xlink"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <g>
                    <rect y="1" width="16" height="2"></rect>
                    <rect y="7" width="16" height="2"></rect>
                    <rect y="13" width="16" height="2"></rect>
                  </g>
                </svg>
              </a>
            </div>
          </div>
        </header>
        <section className="u-clearfix u-section-1" id="sec-b37f">
          <div className="u-clearfix u-sheet u-sheet-1">
            <div className="data-layout-selected u-clearfix u-layout-wrap u-layout-wrap-1">
              <div className="u-layout">
                <div className="u-layout-row">
                  <div className="u-align-left u-container-style u-layout-cell u-left-cell u-size-43 u-layout-cell-1">
                    <div className="u-container-layout u-valign-top u-container-layout-1">
                      <h3 className="u-text u-text-default u-text-1">
                        Dodaj nowy pomysł
                      </h3>
                      <p className="u-text u-text-2">
                        Ma być fajny​ i romantyczny
                        <br />
                      </p>
                    </div>
                  </div>
                  <div className="u-align-left u-container-align-center u-container-style u-layout-cell u-right-cell u-size-17 u-layout-cell-2">
                    <div className="u-container-layout u-valign-middle u-container-layout-2">
                      <a
                        href="#"
                        className="u-align-center u-black u-btn u-button-style u-btn-1"
                      >
                        Dodaj
                      </a>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="data-layout-selected u-clearfix u-layout-wrap u-layout-wrap-2">
              <div className="u-layout">
                <div className="u-layout-row">
                  <div className="u-align-left u-container-style u-layout-cell u-left-cell u-size-43 u-layout-cell-3">
                    <div className="u-container-layout u-container-layout-3">
                      <h3 className="u-text u-text-default u-text-3">
                        Wyświetl dodane pomysły
                      </h3>
                      <p className="u-text u-text-4">Wszystkie są super</p>
                    </div>
                  </div>
                  <div className="u-align-left u-container-align-center u-container-style u-layout-cell u-right-cell u-size-17 u-layout-cell-4">
                    <div className="u-container-layout u-valign-middle u-container-layout-4">
                      <a
                        href="#"
                        className="u-align-center u-black u-btn u-button-style u-btn-2"
                      >
                        Wyświetl
                      </a>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <footer
          className="u-align-center u-clearfix u-footer u-grey-80 u-footer"
          id="sec-07a7"
        >
          <div className="u-clearfix u-sheet u-sheet-1">
            <p className="u-small-text u-text u-text-variant u-text-1">
              Sample text. Click to select the Text Element.
            </p>
          </div>
        </footer>
        <section className="u-backlink u-clearfix u-grey-80">
          <a
            className="u-link"
            href="https://nicepage.com/website-templates"
            target="_blank"
          >
            <span>Website Templates</span>
          </a>
          <p className="u-text">
            <span>created with</span>
          </p>
          <a className="u-link" href="" target="_blank">
            <span>Website Builder Software</span>{" "}
          </a>
          .
        </section>
      </body>
    </html>
  );
};

export default HomePage;
