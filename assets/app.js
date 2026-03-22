
function showSection(id){
  document.querySelectorAll('.section').forEach(s => s.classList.remove('active'));
  document.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
  document.getElementById(id).classList.add('active');
  document.querySelector(`.tab[data-target="${id}"]`).classList.add('active');
  window.scrollTo({top: 0, behavior: 'smooth'});
}

function renderKeywords(targetId, keywords){
  const el = document.getElementById(targetId);
  el.innerHTML = keywords.map(k => `
    <div class="keyword-card">
      <div class="en">${k.en}</div>
      <div class="id">${k.id}</div>
    </div>
  `).join('');
}

function createQuizRenderer(containerId, questions){
  let current = 0, score = 0, answered = false, selected = [];
  const container = document.getElementById(containerId);

  const letters = ['A','B','C','D','E'];

  function renderResult(){
    const pct = Math.round(score / questions.length * 100);
    let msg = pct >= 85 ? '🔥 Keren! Kamu sangat siap UTS.' : pct >= 70 ? '👍 Bagus! Tinggal review bagian yang salah.' : '📚 Ayo ulangi lagi, fondasi kamu masih bisa diperkuat.';
    container.innerHTML = `
      <div class="result-box">
        <div class="result-score">${pct}%</div>
        <div>${score} dari ${questions.length} soal benar</div>
        <p style="margin-top:10px;color:var(--muted)">${msg}</p>
        <div class="btnrow" style="justify-content:center">
          <button class="btn-primary" id="restartQuizBtn">Ulangi Kuis</button>
          <button class="btn-secondary" id="reviewBtn">Review Materi</button>
        </div>
      </div>`;
    document.getElementById('restartQuizBtn').onclick = restart;
    document.getElementById('reviewBtn').onclick = () => showSection('materi');
  }

  function render(){
    if(current >= questions.length) return renderResult();
    const q = questions[current];
    const progress = Math.round(current / questions.length * 100);

    const opts = (q.type === 'cata'
      ? `<div class="checks">` + q.options.map((opt,i)=>`
          <div class="checkopt" data-i="${i}">
            <div class="box"></div><div>${opt}</div>
          </div>`).join('') + `</div>`
      : `<div class="options">` + q.options.map((opt,i)=>`
          <div class="option" data-i="${i}">
            <div class="letter">${letters[i]}</div><div>${opt}</div>
          </div>`).join('') + `</div>`);

    container.innerHTML = `
      <div class="quiz-header">
        <div><span class="quiz-counter">${current+1}</span><span class="quiz-total"> / ${questions.length}</span></div>
        <div class="small">Skor: ${score}</div>
      </div>
      <div class="progress-bar"><div class="progress-fill" style="width:${progress}%"></div></div>
      <div class="question-type-badge">${q.typeName || 'Soal'}</div>
      <div class="question-text">${q.q}</div>
      ${opts}
      <div class="btnrow">
        <button class="btn-primary" id="checkBtn" disabled>Periksa Jawaban</button>
      </div>
      <div class="explanation" id="explanationBox"></div>
      <div class="btnrow" id="nextRow" style="display:none">
        <button class="btn-primary" id="nextBtn">${current < questions.length-1 ? 'Soal Berikutnya →' : 'Lihat Hasil 🎉'}</button>
      </div>
    `;

    selected = [];
    answered = false;

    if(q.type === 'cata'){
      container.querySelectorAll('.checkopt').forEach(el=>{
        el.onclick = ()=>{
          if(answered) return;
          const i = Number(el.dataset.i);
          if(selected.includes(i)){
            selected = selected.filter(x => x !== i);
            el.classList.remove('checked');
            el.querySelector('.box').textContent = '';
          } else {
            selected.push(i);
            el.classList.add('checked');
            el.querySelector('.box').textContent = '✓';
          }
          document.getElementById('checkBtn').disabled = selected.length === 0;
        };
      });
    } else {
      container.querySelectorAll('.option').forEach(el=>{
        el.onclick = ()=>{
          if(answered) return;
          container.querySelectorAll('.option').forEach(o=>o.classList.remove('selected'));
          el.classList.add('selected');
          selected = [Number(el.dataset.i)];
          document.getElementById('checkBtn').disabled = false;
        };
      });
    }

    document.getElementById('checkBtn').onclick = check;
    document.getElementById('nextBtn') && (document.getElementById('nextBtn').onclick = next);
  }

  function check(){
    if(answered || selected.length === 0) return;
    answered = true;
    const q = questions[current];
    const exp = document.getElementById('explanationBox');
    let correct = false;

    if(q.type === 'cata'){
      const ans = new Set(q.answers);
      const sel = new Set(selected);
      correct = [...ans].every(x => sel.has(x)) && [...sel].every(x => ans.has(x));

      container.querySelectorAll('.checkopt').forEach(el=>{
        const i = Number(el.dataset.i);
        if(ans.has(i)) el.classList.add('correct');
        if(sel.has(i) && !ans.has(i)) el.classList.add('wrong');
      });
    } else {
      correct = selected[0] === q.answer;
      container.querySelectorAll('.option').forEach(el=>{
        const i = Number(el.dataset.i);
        if(i === q.answer) el.classList.add('correct');
        if(i === selected[0] && i !== q.answer) el.classList.add('wrong');
      });
    }

    if(correct) score++;
    exp.innerHTML = `<h4>${correct ? '✅ Benar' : '❌ Belum tepat'}</h4><div>${q.explanation}</div>`;
    exp.classList.add('show');
    document.getElementById('checkBtn').disabled = true;
    document.getElementById('nextRow').style.display = 'flex';
  }

  function next(){ current++; render(); }
  function restart(){ current = 0; score = 0; answered = false; selected = []; render(); }
  render();
  return { restart };
}
