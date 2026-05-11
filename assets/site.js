// ============================================================
// Go Learning - Shared Site JS
// Initializes Prism syntax highlighting and Mermaid diagrams.
// ============================================================

(function () {
  function init() {
    // Mark <pre> blocks with language label
    document.querySelectorAll('pre[class*="language-"]').forEach((pre) => {
      const cls = pre.className.match(/language-([\w-]+)/);
      if (cls && !pre.hasAttribute('data-lang')) {
        pre.setAttribute('data-lang', cls[1]);
      }
    });

    // Init mermaid (uses dark theme matching Catppuccin)
    if (window.mermaid) {
      window.mermaid.initialize({
        startOnLoad: true,
        theme: 'dark',
        themeVariables: {
          darkMode: true,
          background: '#181825',
          primaryColor: '#313244',
          primaryTextColor: '#cdd6f4',
          primaryBorderColor: '#cba6f7',
          lineColor: '#89b4fa',
          secondaryColor: '#45475a',
          tertiaryColor: '#11111b',
          fontFamily: 'ui-monospace, "JetBrains Mono", monospace',
          fontSize: '14px',
          nodeBorder: '#cba6f7',
          mainBkg: '#313244',
          clusterBkg: '#1e1e2e',
          clusterBorder: '#cba6f7',
          actorBkg: '#313244',
          actorBorder: '#cba6f7',
          actorTextColor: '#cdd6f4',
          actorLineColor: '#89b4fa',
          signalColor: '#cdd6f4',
          signalTextColor: '#cdd6f4',
          labelBoxBkgColor: '#45475a',
          labelBoxBorderColor: '#cba6f7',
          labelTextColor: '#cdd6f4',
          loopTextColor: '#cdd6f4',
          noteBkgColor: '#f9e2af',
          noteTextColor: '#1e1e2e',
          activationBkgColor: '#94e2d5',
          sequenceNumberColor: '#1e1e2e',
        },
        flowchart: { curve: 'basis', useMaxWidth: true, htmlLabels: true },
        sequence: { useMaxWidth: true, mirrorActors: false, boxMargin: 10 },
      });
    }

    // Highlight current sidebar entry by matching pathname
    const path = window.location.pathname;
    document.querySelectorAll('.sidebar a').forEach((a) => {
      if (a.getAttribute('href') && path.endsWith(a.getAttribute('href').replace(/^\.\.?\//, ''))) {
        a.classList.add('active');
      }
    });

    // Build auto-TOC from H2 in content if a .toc-auto placeholder is present
    document.querySelectorAll('.toc-auto').forEach((tocEl) => {
      const main = document.querySelector('main') || document.body;
      const headings = main.querySelectorAll('h2, h3');
      if (!headings.length) return;
      const ul = document.createElement('ul');
      headings.forEach((h) => {
        if (!h.id) {
          h.id = h.textContent
            .toLowerCase()
            .replace(/[^\w一-龥]+/g, '-')
            .replace(/^-|-$/g, '');
        }
        const li = document.createElement('li');
        li.style.marginLeft = h.tagName === 'H3' ? '1em' : '0';
        const a = document.createElement('a');
        a.href = '#' + h.id;
        a.textContent = h.textContent;
        li.appendChild(a);
        ul.appendChild(li);
      });
      const title = document.createElement('div');
      title.className = 'toc-title';
      title.textContent = '本页导航';
      tocEl.innerHTML = '';
      tocEl.appendChild(title);
      tocEl.appendChild(ul);
    });
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
