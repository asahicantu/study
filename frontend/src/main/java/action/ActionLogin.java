package action;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;
import metier.data.Client;
import metier.data.Personne;
import metier.service.Service;

public class ActionLogin extends Action {

    @Override
    public boolean executer(HttpServletRequest request) {
        Integer login = Integer.parseInt(request.getParameter("login"));
        String password = request.getParameter("password");

        Personne personne = Service.chercherPersonne(login, password);

        if (personne != null) {
            request.setAttribute("statut", Boolean.TRUE);
            HttpSession session = request.getSession();
            session.setAttribute("personne", personne);
            if (personne instanceof Client) {
                request.setAttribute("personne", "client");
            } else {
                request.setAttribute("personne", "employe");
            }
        } else {
            request.setAttribute("statut", Boolean.FALSE);
        }
        return true;
    }
}
