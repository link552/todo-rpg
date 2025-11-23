document.addEventListener('DOMContentLoaded', _ => {
    if (LEVEL_PROGRESS != null) {

        var progressBefore = document.getElementById('progress-before');
        var progressBar = document.getElementById('progress-bar');
        var progressAfter = document.getElementById('progress-after');

        // The first record always sets the initial state.
        var init = LEVEL_PROGRESS[0];
        progressBefore.innerText = init.FromLevel;
        progressAfter.innerText = init.ToLevel;
        progressBar.style.width = init.ToPercent + '%';

        var interval = null;
        var modalEl = document.getElementById('progress-dialog');
        var modal = new bootstrap.Modal(modalEl);
        modal.show();
        modalEl.addEventListener('hidden.bs.modal', _ => {
            if (interval != null) {
                clearInterval(interval);
            }
        });

        var confetti = new JSConfetti(document.getElementById('canvas'));
        var i = 1;

        // Iterate through the rest of the level progress on a timer.
        interval = setInterval(_ => {
            if (i == LEVEL_PROGRESS.length) {
                clearInterval(interval);
                return;
            }

            var p = LEVEL_PROGRESS[i];

            if (p.ToPercent == 0) {
                confetti.addConfetti();
                progressBar.style.transition = 'none';
            } else {
                progressBar.style.transition = 'width 0.6s ease';
            }

            progressBefore.innerText = p.FromLevel;
            progressAfter.innerText = p.ToLevel;
            progressBar.style.width = p.ToPercent + '%';

            i++;
        }, 600);
    }
});
